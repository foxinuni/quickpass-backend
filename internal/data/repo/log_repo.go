package repo

import (
	"context"
	"time"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
	"github.com/foxinuni/quickpass-backend/internal/data/stores"
)

type LogRepository interface {
	NewAction(occasionID int) (bool, error)
}

type StoreLogRepository struct {
	logStore stores.LogStore
}

func NewStoreLogRepository(logStore stores.LogStore) LogRepository {
	return &StoreLogRepository{
		logStore: logStore,
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
