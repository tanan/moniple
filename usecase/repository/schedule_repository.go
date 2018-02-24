package repository

import "github.com/tanan/moniple/domain"

type ScheduleRepository interface {
	FindAll() ([]domain.Schedule, error)
	FindById(id domain.ScheduleId) (domain.Schedule, error)
	Store(schedule domain.Schedule) error
}
