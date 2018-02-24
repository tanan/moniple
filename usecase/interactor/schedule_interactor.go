package interactor

import (
	"github.com/tanan/moniple/usecase/repository"
	"github.com/tanan/moniple/domain"
)

type ScheduleInteractor struct {
	Repository repository.ScheduleRepository
}

func (i *ScheduleInteractor) Store(schedule domain.Schedule) error {
	return i.Repository.Store(schedule)
}

func (i *ScheduleInteractor) FindScheduleById(id domain.ScheduleId) (domain.Schedule, error) {
	return i.Repository.FindById(id)
}

func (i *ScheduleInteractor) FindAll() ([]domain.Schedule, error) {
	return i.Repository.FindAll()
}