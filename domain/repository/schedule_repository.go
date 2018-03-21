package repository

import (
	"github.com/tanan/moniple/domain/model"
)

type ScheduleRepository interface {
	FindAll() ([]model.Schedule, error)
	FindById(id model.ScheduleID) (model.Schedule, error)
	Store(schedule model.Schedule) error
}
