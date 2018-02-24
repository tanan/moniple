package database

import (
	"github.com/tanan/moniple/domain"
	"github.com/tanan/moniple/infrastracture/database"
	"github.com/tanan/moniple/interfaces/database/model"
)

type ScheduleRepository struct {
	SqlHandler
}

func NewScheduleRepository(connInfo string) ScheduleRepository {
	return ScheduleRepository{
		database.NewSqlHandler(connInfo),
	}
}

func (r ScheduleRepository) FindAll() ([]domain.Schedule, error) {
	return []domain.Schedule{}, nil
}

func (r ScheduleRepository) FindById(id domain.ScheduleId) (domain.Schedule, error) {
	return domain.Schedule{}, nil
}

func (r ScheduleRepository) Store(schedule domain.Schedule) error {
	s := model.Schedule{
		Id:       int64(schedule.Id),
		Name:     schedule.Name,
		URL:      schedule.URL,
		Interval: schedule.Interval,
		Status:   r.getStatus(schedule.Active),
	}
	if r.NewRecord(s) {
		return r.Create(s)
	}
	return r.Save(s)
}

func (r ScheduleRepository) getStatus(active bool) string {
	if active {
		return "active"
	}
	return "inactive"
}
