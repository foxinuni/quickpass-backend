package repo

import (
	"context"
	"time"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
	"github.com/foxinuni/quickpass-backend/internal/data/stores"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
)

type LogRepository interface {
	NewAction(occasionID int) (bool, error)
	GetLogs(eventId *int, bookingId *int) ([]*entities.LogHistory, error)
}

type StoreLogRepository struct {
	logStore stores.LogStore
	occasionStore stores.OccasionStore
	userStore stores.UserStore
}

func NewStoreLogRepository(logStore stores.LogStore, occasionStore stores.OccasionStore, userStore stores.UserStore) LogRepository {
	return &StoreLogRepository{
		logStore: logStore,
		occasionStore: occasionStore,
		userStore: userStore,
	}
}

func (r *StoreLogRepository) NewAction(occasionID int) (bool, error) {
	log, err := r.logStore.GetLastFromOcassion(context.Background(), occasionID)
	if err != nil {
		return false, err
	}

	var isInside bool = true
	if log != nil {
		//if last log ijsn't nil, then the value of this new log
		//is the contrast of the previously last
		//by default inside is true since if it's the first time it must have been entering
		isInside = !log.IsInside
	}

	var newLog models.Log = models.Log{
		OccasionID: occasionID,
		IsInside:   isInside,
		Time:       time.Now(),
	}

	err = r.logStore.Create(context.Background(), &newLog)
	if err != nil {
		return false, nil
	}
	return isInside, nil

}

func (r * StoreLogRepository) GetLogs(eventId *int, bookingId *int) ([]*entities.LogHistory, error){
	ocassions, err:=r.occasionStore.GetAll(context.Background(), stores.OccasionFilter{
		EventID: eventId,
		BookingID: bookingId,
	})

	if err != nil {
		return nil, err
	}

	allLogs := make([]*entities.LogHistory, 0)

	for _, occasion := range  ocassions{
		user, err := r.userStore.GetById(context.Background(), occasion.UserID)
		if err == nil{
			logs, err :=r.logStore.GetAll(context.Background(), stores.LogFilter{
				OccasionId: &occasion.OccasionID,
			})
			
			if err == nil{
				for _, log := range logs{
					allLogs = append(allLogs, entities.NewLogHistory(log.LogID, user.Email, log.IsInside, log.Time))
				}
			}
		}
	}
	return allLogs, nil

}
