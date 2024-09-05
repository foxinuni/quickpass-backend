package services

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
	"github.com/rs/zerolog/log"
	"github.com/xuri/excelize/v2"
)

type ImportService interface {
	ImportFromFile(reader io.Reader) error
}

type ExcelImportService struct {
	// userRepo  repo.UserRepository
	// eventRepo repo.EventRepository
}

func NewExcelImportService() *ExcelImportService {
	return &ExcelImportService{}
}

func (s *ExcelImportService) ImportFromFile(reader io.Reader) error {
	f, err := excelize.OpenReader(reader)
	if err != nil {
		return err
	}
	defer f.Close()

	// Get all the rows for the first sheet
	fmt.Printf("Sheet names: %v\n", f.GetSheetName(0))
	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		return err
	}

	// email, number (client)
	// name, start_date, end_date, address (event)
	// address, entry_date, leaving_date, is_house (house)

	// Iterate over the rows
	for i, row := range rows {
		if i < 2 {
			continue
		}

		var user *entities.User
		var event *entities.Event
		var accomodation *entities.Accomodation
		var booking *entities.Booking

		fmt.Printf("row: %+v\n", row)

		// email, number (client)
		email, number := row[0], row[1]
		user = entities.NewUser(0, email, number)

		// name, start_date, end_date, address (event)
		name := row[2]
		if name != "" {
			startDate, err := time.Parse("2006-01-02", row[3])
			if err != nil {
				log.Warn().Err(err).Msgf("failed to parse date: %s when scanning for event", row[3])
				return err
			}

			endDate, err := time.Parse("2006-01-02", row[4])
			if err != nil {
				log.Warn().Err(err).Msgf("failed to parse date: %s when scanning for event", row[4])
				return err
			}

			address := row[5]

			event = entities.NewEvent(0, name, address, startDate, endDate)
		}

		// address, entry_date, leaving_date, is_house (booking)
		address := row[6]
		if address != "" {
			entryDate, err := time.Parse("2006-01-02", row[7])
			if err != nil {
				log.Warn().Err(err).Msgf("failed to parse date: %s when scanning for house", row[7])
				return err
			}

			leavingDate, err := time.Parse("2006-01-02", row[8])
			if err != nil {
				log.Warn().Err(err).Msgf("failed to parse date: %s when scanning for house", row[8])
				return err
			}

			isHouse, err := strconv.ParseBool(row[9])
			if err != nil {
				log.Warn().Err(err).Msgf("failed to parse bool: %s when scanning for house", row[9])
				return err
			}

			// Create the accomodation
			accomodation = entities.NewAccomodation(0, isHouse, address)

			// Create the booking
			booking = entities.NewBooking(0, nil, entryDate, leavingDate)
		}

		// Debug print all the data
		fmt.Printf("user: %v\n", user)
		fmt.Printf("event: %v\n", event)
		fmt.Printf("accomodation: %v\n", accomodation)
		fmt.Printf("booking: %v\n", booking)

		/*
			// Find user if it exists
			if found, err := s.userRepo.GetByEmail(user.Email); err != nil {
				if err == stores.ErrUserNotFound {
					if err := s.userRepo.Create(user); err != nil {
						log.Warn().Err(err).Msgf("failed to create user with email: %s", user.Email)
						return err
					}
				} else {
					log.Warn().Err(err).Msgf("failed to find user with email: %s", user.Email)
					return err
				}
			} else {
				user = found
			}

			// Create the event
			if event != nil {
				if err := s.eventRepo.Create(event); err != nil {
					log.Warn().Err(err).Msgf("failed to create event with name: %s", event.Name)
					return err
				}
			}
		*/
	}
	return nil
}

func (s *ExcelImportService) GetOrCreateUserFromRow(file *excelize.File, sheet string, row string) (*entities.User, error) {
	// Get the data from the row
	return nil, nil
}
