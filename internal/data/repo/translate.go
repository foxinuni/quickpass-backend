package repo

import (
	"github.com/foxinuni/quickpass-backend/internal/data/models"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
)

func ModelToAccomodation(accomodation *models.Accomodation) *entities.Accomodation {
	return entities.NewAccomodation(accomodation.AccomodationID, accomodation.IsHouse, accomodation.Address)
}

func AccomodationToModel(accomodation *entities.Accomodation) *models.Accomodation {
	return models.NewAccomodation(accomodation.GetAccomodationID(), accomodation.GetIsHouse(), accomodation.GetAddress())
}

func ModelToBooking(model *models.Booking, accomodation *entities.Accomodation) *entities.Booking {
	return entities.NewBooking(model.BookingID, accomodation, model.EntryDate, model.ExitDate)
}

func BookingToModel(booking *entities.Booking) *models.Booking {
	return models.NewBooking(booking.GetBookingID(), booking.GetAccomodation().GetAccomodationID(), booking.GetEntryDate(), booking.GetExitDate())
}

func ModelToEvent(model *models.Event) *entities.Event {
	return entities.NewEvent(model.EventID, model.Name, model.Address, model.StartDate, model.EndDate)
}

func EventToModel(event *entities.Event) *models.Event {
	return models.NewEvent(event.GetEventID(), event.GetName(), event.GetAddress(), event.GetStartDate(), event.GetEndDate())
}

func ModelToLog(model *models.Log, occasion *entities.Occasion) *entities.Log {
	return entities.NewLog(model.LogID, occasion, model.IsInside, model.Time)
}

func LogToModel(log *entities.Log) *models.Log {
	return models.NewLog(log.GetLogID(), log.GetOccasion().GetOccasionID(), log.GetIsInside(), log.GetTime())
}

func ModelToOccasion(occasions *models.Occasion, user *entities.User, event *entities.Event, booking *entities.Booking, state *entities.State, lastLog *models.Log) *entities.Occasion {
	var isInside bool = false
	//is inside by default will be false, if there are no logs, then by default its outside, if there's a log, then we ask it if its inside or outside
	if lastLog != nil {
		isInside = lastLog.IsInside
	}
	return entities.NewOccasion(occasions.OccasionID, user, event, booking, state, isInside)
}

func OccasionToModel(occasion *entities.Occasion) *models.Occasion {
	var eventID *int = nil
	var bookingID *int = nil

	if occasion.GetEvent() != nil{
		var id int =  occasion.GetEvent().GetEventID()
		eventID = &id
	}
	if occasion.GetBooking() != nil{
		var id int =  occasion.GetBooking().GetBookingID()
		bookingID = &id
	}
	return models.NewOccasion(
		occasion.GetOccasionID(),
		occasion.GetUser().GetUserID(),
		eventID,
		bookingID,
		occasion.GetState().GetStateID(),
	)
}

func SessionToModel(session *entities.Session) *models.Session {
	return models.NewSession(session.GetSessionID(), session.GetUser().GetUserID(), session.GetEnabled(), session.GetToken(), session.GetPhoneModel(), session.GetIMEI())
}

func ModelToSession(session *models.Session, user *entities.User) *entities.Session {
	return entities.NewSession(session.SessionID, user, session.Enabled, session.Token, session.PhoneModel, session.IMEI)
}

func ModelToState(model *models.State) *entities.State {
	return entities.NewState(model.StateID, model.StateName)
}

func StateToModel(state *entities.State) *models.State {
	return models.NewState(state.GetStateID(), state.GetStateName())
}

func UserToModel(user *entities.User) *models.User {
	return models.NewUser(user.GetUserID(), user.GetEmail(), user.GetNumber())
}

func ModelToUser(user *models.User) *entities.User {
	return entities.NewUser(user.UserID, user.Email, user.Number)
}
