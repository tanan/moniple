package usecase

import (
	"github.com/tanan/moniple/domain/model"
	"github.com/tanan/moniple/domain/repository"
)

type ScheduleUseCase interface {
	Store(schedule model.Schedule) error
	FindScheduleById(id model.ScheduleID) (model.Schedule, error)
	FindAll() ([]model.Schedule, error)
}

type scheduleUseCase struct {
	Repository repository.ScheduleRepository
}

func NewScheduleUseCase(repository repository.ScheduleRepository) ScheduleUseCase {
	return &scheduleUseCase{
		Repository: repository,
	}
}

func (i *scheduleUseCase) Store(schedule model.Schedule) error {
	return i.Repository.Store(schedule)
}

func (i *scheduleUseCase) FindScheduleById(id model.ScheduleID) (model.Schedule, error) {
	return i.Repository.FindById(id)
}

func (i *scheduleUseCase) FindAll() ([]model.Schedule, error) {
	return i.Repository.FindAll()
}
