package registry

import (
	"github.com/tanan/moniple/application"
	"github.com/tanan/moniple/domain/repository"
	"github.com/tanan/moniple/infrastracture/persistence/datastore"
	"github.com/tanan/moniple/interfaces/api/handler"
)

type Registry struct {
	useCase    usecase.ScheduleUseCase
	repository repository.ScheduleRepository
	Handler    handler.ScheduleHandler
}

func NewRegistry(sqlHandler *datastore.SQLHandler) *Registry {
	repository := NewScheduleRepository(sqlHandler)
	usecase := NewScheduleUseCase(repository)
	return &Registry{
		repository: repository,
		useCase:    usecase,
		Handler:    handler.NewScheduleHandler(usecase),
	}
}

func NewScheduleUseCase(repository repository.ScheduleRepository) usecase.ScheduleUseCase {
	return usecase.NewScheduleUseCase(repository)
}

func NewScheduleRepository(handler *datastore.SQLHandler) repository.ScheduleRepository {
	return datastore.NewScheduleRepository(handler)
}
