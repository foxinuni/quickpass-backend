package repo

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
	"github.com/foxinuni/quickpass-backend/internal/data/stores"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
)

type OccasionLookupFilterOption func(*OccasionLookupFilter)

type OccasionLookupFilter struct {
	User           *entities.User
	Event          *entities.Event
	Booking        *entities.Booking
	State          *entities.State
	TypeOfOccasion *bool
}

func DefaultOccasionLookupFilter() *OccasionLookupFilter {
	return &OccasionLookupFilter{}
}

func OccasionForUser(user *entities.User) OccasionLookupFilterOption {
	return func(f *OccasionLookupFilter) {
		f.User = user
	}
}

func OccasionForEvent(event *entities.Event) OccasionLookupFilterOption {
	return func(f *OccasionLookupFilter) {
		f.Event = event
	}
}

func OccasionForBooking(booking *entities.Booking) OccasionLookupFilterOption {
	return func(f *OccasionLookupFilter) {
		f.Booking = booking
	}
}

func OccasionForState(state *entities.State) OccasionLookupFilterOption {
	return func(f *OccasionLookupFilter) {
		f.State = state
	}
}

func OccasionForType(typeOfOccasion bool) OccasionLookupFilterOption {
	return func(f *OccasionLookupFilter) {
		f.TypeOfOccasion = &typeOfOccasion
	}
}

func LookupToFilter(lookup *OccasionLookupFilter) stores.OccasionFilter {
	var occasionFilter stores.OccasionFilter
	if lookup.User != nil {
		occasionFilter.UserID = &lookup.User.UserID
	}

	if lookup.Event != nil {
		occasionFilter.EventID = &lookup.Event.EventID
	}

	if lookup.Booking != nil {
		occasionFilter.BookingID = &lookup.Booking.BookingID
	}

	if lookup.State != nil {
		occasionFilter.StateID = &lookup.State.StateID
	}

	if lookup.TypeOfOccasion != nil {
		occasionFilter.TypeOccasion = lookup.TypeOfOccasion
	}

	return occasionFilter
}

type OccasionRepository interface {
	GetAll(filters ...OccasionLookupFilterOption) ([]*entities.Occasion, error)
	GetById(occasionID int) (*entities.Occasion, error)
	Create(occasion *entities.Occasion) error
	Update(occasion *entities.Occasion) error
	Delete(occasionID int) error
}

type StoreOccasionRepository struct {
	occasionStore     stores.OccasionStore
	userStore         stores.UserStore
	eventStore        stores.EventStore
	bookingStore      stores.BookingStore
	accomodationStore stores.AccomodationStore
	stateStore        stores.StateStore
	logStore          stores.LogStore
}

func NewStoreOccasionRepository(
	occasionStore stores.OccasionStore,
	userStore stores.UserStore,
	eventStore stores.EventStore,
	bookingStore stores.BookingStore,
	accomodationStore stores.AccomodationStore,
	stateStore stores.StateStore,
) OccasionRepository {
	return &StoreOccasionRepository{
		occasionStore:     occasionStore,
		userStore:         userStore,
		eventStore:        eventStore,
		bookingStore:      bookingStore,
		accomodationStore: accomodationStore,
		stateStore:        stateStore,
	}
}

func (r *StoreOccasionRepository) PopulateOccasion(occasion *models.Occasion) (*entities.Occasion, error) {
	// Get the user from the store
	user, err := r.userStore.GetById(context.Background(), occasion.UserID)
	if err != nil {
		return nil, err
	}

	// Get the event from the store
	event, err := r.eventStore.GetById(context.Background(), occasion.EventID)
	if err != nil {
		return nil, err
	}

	// Get the booking from the store
	booking, err := r.bookingStore.GetById(context.Background(), occasion.BookingID)
	if err != nil {
		return nil, err
	}

	// Get the accomodation from the store
	accomodation, err := r.accomodationStore.GetById(context.Background(), booking.AccomodationID)
	if err != nil {
		return nil, err
	}

	// Get the state from the store
	state, err := r.stateStore.GetById(context.Background(), occasion.StateID)
	if err != nil {
		return nil, err
	}

	//get the last log from store
	log, err := r.logStore.GetLastFromOcassion(context.Background(), occasion.OccasionID)
	if err != nil {
		return nil, err
	}

	// Convert the result to a Occasion entity
	return ModelToOccasion(
		occasion,
		ModelToUser(user),
		ModelToEvent(event),
		ModelToBooking(booking, ModelToAccomodation(accomodation)),
		ModelToState(state),
		log.IsInside,
	), nil
}

func (r *StoreOccasionRepository) GetAll(filters ...OccasionLookupFilterOption) ([]*entities.Occasion, error) {
	lookup := DefaultOccasionLookupFilter()
	for _, f := range filters {
		f(lookup)
	}

	// Get the occasion from the store
	filter := LookupToFilter(lookup)
	occasions, err := r.occasionStore.GetAll(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	// Populate the occasions
	var result []*entities.Occasion
	for _, occasion := range occasions {
		populated, err := r.PopulateOccasion(&occasion)
		if err != nil {
			return nil, err
		}

		result = append(result, populated)
	}

	return result, nil
}

func (r *StoreOccasionRepository) GetById(occasionID int) (*entities.Occasion, error) {
	// Get the occasion from the store
	occasion, err := r.occasionStore.GetById(context.Background(), occasionID)
	if err != nil {
		return nil, err
	}

	// Populate the occasion
	return r.PopulateOccasion(occasion)
}

func (r *StoreOccasionRepository) Create(occasion *entities.Occasion) error {
	// Convert the occasion to a model
	model := OccasionToModel(occasion)

	// Create the occasion in the store
	if err := r.occasionStore.Create(context.Background(), model); err != nil {
		return err
	}

	temp, err := r.PopulateOccasion(model)
	if err != nil {
		return err
	}

	*occasion = *temp
	return nil
}

func (r *StoreOccasionRepository) Update(occasion *entities.Occasion) error {
	// Convert the occasion to a model
	model := OccasionToModel(occasion)

	// Update the occasion in the store
	if err := r.occasionStore.Update(context.Background(), model); err != nil {
		return err
	}

	// Populate the occasion
	temp, err := r.PopulateOccasion(model)
	if err != nil {
		return err
	}

	*occasion = *temp
	return nil
}

func (r *StoreOccasionRepository) Delete(occasionID int) error {
	// Delete the occasion from the store
	return r.occasionStore.Delete(context.Background(), occasionID)
}
