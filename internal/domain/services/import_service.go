package services

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/foxinuni/quickpass-backend/internal/data/repo"
	"github.com/foxinuni/quickpass-backend/internal/data/stores"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
	"github.com/rs/zerolog/log"
	"github.com/xuri/excelize/v2"
)

type ImportService interface {
	ImportFromFile(reader io.Reader) (int, error)
}

type ExcelImportService struct {
	userRepo     repo.UserRepository
	eventRepo    repo.EventRepository
	accomRepo    repo.AccomodationRepository
	bookingRepo  repo.BookingRepository
	occasionRepo repo.OccasionRepository
	stateService StateService
}

func NewExcelImportService(
	userRepo repo.UserRepository,
	event repo.EventRepository,
	accomRepo repo.AccomodationRepository,
	bookingRepo repo.BookingRepository,
	occasionRepo repo.OccasionRepository,
	stateService StateService,
) ImportService {
	return &ExcelImportService{
		userRepo:     userRepo,
		eventRepo:    event,
		accomRepo:    accomRepo,
		bookingRepo:  bookingRepo,
		occasionRepo: occasionRepo,
		stateService: stateService,
	}
}

func (s *ExcelImportService) ImportFromFile(reader io.Reader) (int, error) {
	f, err := excelize.OpenReader(reader)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	// Get all the rows for the first sheet
	row := 3
	counter := 0
	sheet := f.GetSheetName(0)

	// Get or create state
	state, err := s.stateService.GetOrCreateState(StateRegistered)
	if err != nil {
		return 0, err
	}

	for {
		user, err := s.parseUser(f, sheet, row)
		if err != nil {
			if err == ErrEmailEmpty {
				break
			}
		}

		// Get or create user
		if err := s.getOrCreateUser(user); err != nil {
			return 0, err
		}

		event, err := s.parseEvent(f, sheet, row)
		if err != nil && err != ErrNameEmpty {
			return 0, err
		}

		// Get or create event
		if event != nil {
			if err := s.getOrCreateEvent(event); err != nil {
				return 0, err
			}
		}

		accomodation, booking, err := s.parseBooking(f, sheet, row)
		if err != nil && err != ErrAccAddressEmpty {
			return 0, err
		}

		// Get or create booking
		if accomodation != nil && booking != nil {
			if err := s.getOrCreateBooking(accomodation, booking); err != nil {
				return 0, err
			}
		}

		// Create occasion
		occasion := entities.NewOccasion(0, user, event, booking, state, false)
		if err := s.occasionRepo.Create(occasion); err != nil {
			return 0, err
		}

		log.Debug().Msgf("Imported (user: %v, event: %v, booking: %v)", user, event, booking)
		row++
		counter++
	}

	return counter, nil
}

var ErrEmailEmpty = fmt.Errorf("email is empty")
var ErrNumberEmpty = fmt.Errorf("number is empty")

func (s *ExcelImportService) parseUser(file *excelize.File, sheet string, row int) (*entities.User, error) {
	// Read email and number from the row
	email, err := file.GetCellValue(sheet, fmt.Sprintf("A%d", row))
	if err != nil {
		return nil, err
	}

	if email == "" {
		return nil, ErrEmailEmpty
	}

	// Read the number from the row
	number, err := file.GetCellValue(sheet, fmt.Sprintf("B%d", row))
	if err != nil {
		return nil, err
	}

	if number == "" {
		return nil, ErrNumberEmpty
	}

	// Create the user
	user := entities.NewUser(0, email, number)
	return user, nil
}

func (s *ExcelImportService) getOrCreateUser(entity *entities.User) error {
	user, err := s.userRepo.GetByEmail(entity.Email)
	if err != nil {
		if err == stores.ErrUserNotFound {
			if err := s.userRepo.Create(entity); err != nil {
				return err
			}

			return nil
		} else {
			return err
		}
	}

	*entity = *user
	return nil
}

var ErrNameEmpty = fmt.Errorf("name is empty")
var ErrStartDateEmpty = fmt.Errorf("start date is empty")
var ErrEndDateEmpty = fmt.Errorf("end date is empty")
var ErrAddressEmpty = fmt.Errorf("address is empty")

func (s *ExcelImportService) parseEvent(file *excelize.File, sheet string, row int) (*entities.Event, error) {
	// Read the name from the row
	name, err := file.GetCellValue(sheet, fmt.Sprintf("C%d", row))
	if err != nil {
		return nil, err
	}

	if name == "" {
		return nil, ErrNameEmpty
	}

	// Read the start date from the row
	startDateStr, err := file.GetCellValue(sheet, fmt.Sprintf("D%d", row))
	if err != nil {
		return nil, err
	}

	if startDateStr == "" {
		return nil, ErrStartDateEmpty
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return nil, err
	}

	// Read the end date from the row
	endDateStr, err := file.GetCellValue(sheet, fmt.Sprintf("E%d", row))
	if err != nil {
		return nil, err
	}

	if endDateStr == "" {
		return nil, ErrEndDateEmpty
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return nil, err
	}

	// Read the address from the row
	address, err := file.GetCellValue(sheet, fmt.Sprintf("F%d", row))
	if err != nil {
		return nil, err
	}

	// Create the event
	event := entities.NewEvent(0, name, address, startDate, endDate)
	return event, nil
}

func (s *ExcelImportService) getOrCreateEvent(entity *entities.Event) error {
	event, err := s.eventRepo.GetByName(entity.Name)
	if err != nil {
		if err == stores.ErrEventNotFound {
			if err := s.eventRepo.Create(entity); err != nil {
				return err
			}

			return nil
		} else {
			return err
		}
	}

	*entity = *event
	return nil
}

var ErrAccAddressEmpty = fmt.Errorf("address is empty")
var ErrEntryDateEmpty = fmt.Errorf("entry date is empty")
var ErrLeavingDateEmpty = fmt.Errorf("leaving date is empty")
var ErrIsHouseEmpty = fmt.Errorf("is house is empty")

func (s *ExcelImportService) parseBooking(file *excelize.File, sheet string, row int) (*entities.Accomodation, *entities.Booking, error) {
	// Read the address from the row
	address, err := file.GetCellValue(sheet, fmt.Sprintf("G%d", row))
	if err != nil {
		return nil, nil, err
	}

	if address == "" {
		return nil, nil, ErrAccAddressEmpty
	}

	// Read the entry date from the row
	entryDateStr, err := file.GetCellValue(sheet, fmt.Sprintf("H%d", row))
	if err != nil {
		return nil, nil, err
	}

	if entryDateStr == "" {
		return nil, nil, ErrEntryDateEmpty
	}

	entryDate, err := time.Parse("2006-01-02", entryDateStr)
	if err != nil {
		return nil, nil, err
	}

	// Read the leaving date from the row
	leavingDateStr, err := file.GetCellValue(sheet, fmt.Sprintf("I%d", row))
	if err != nil {
		return nil, nil, err
	}

	if leavingDateStr == "" {
		return nil, nil, ErrLeavingDateEmpty
	}

	leavingDate, err := time.Parse("2006-01-02", leavingDateStr)
	if err != nil {
		return nil, nil, err
	}

	// Read the is house from the row
	isHouseStr, err := file.GetCellValue(sheet, fmt.Sprintf("J%d", row))
	if err != nil {
		return nil, nil, err
	}

	if isHouseStr == "" {
		return nil, nil, ErrIsHouseEmpty
	}

	isHouse, err := strconv.ParseBool(isHouseStr)
	if err != nil {
		return nil, nil, err
	}

	accomodation := entities.NewAccomodation(0, isHouse, address)
	booking := entities.NewBooking(0, nil, entryDate, leavingDate)

	return accomodation, booking, nil
}

func (s *ExcelImportService) getOrCreateBooking(accomodation *entities.Accomodation, booking *entities.Booking) error {
	accom, err := s.accomRepo.GetByAddress(accomodation.Address)
	if err != nil {
		if err == stores.ErrAccomodationNotFound {
			if err := s.accomRepo.Create(accomodation); err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		*accomodation = *accom
	}

	// Update the booking with the accomodation
	booking.Accomodation = accomodation

	// Create booking
	return s.bookingRepo.Create(booking)
}
