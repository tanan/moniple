package datastore

import (
	"github.com/tanan/moniple/domain/model"
	"github.com/tanan/moniple/domain/repository"
	"github.com/tanan/moniple/infrastracture/persistence/datastore/schema"
)

type scheduleRepository struct {
	*SQLHandler
}

func NewScheduleRepository(handler *SQLHandler) repository.ScheduleRepository {
	return &scheduleRepository{
		handler,
	}
}

func (r *scheduleRepository) FindAll() ([]model.Schedule, error) {
	var m []model.Schedule
	err := r.Find(&m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *scheduleRepository) Store(schedule model.Schedule) error {
	s := &schema.Schedule{
		Id:       int64(schedule.ID),
		Name:     schedule.Name,
		URL:      schedule.URL,
		Method:   schedule.Method,
		Interval: schedule.Interval,
		Status:   r.getStatus(schedule.Active),
	}

	if r.NewRecord(s) {
		return r.Create(s)
	}
	return r.Save(s)
}

func (r *scheduleRepository) FindById(id model.ScheduleID) (model.Schedule, error) {
	return model.Schedule{}, nil
}

func (r scheduleRepository) getStatus(active bool) string {
	if active {
		return "active"
	}
	return "inactive"
}
