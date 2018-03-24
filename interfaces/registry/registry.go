package registry

import (
	"github.com/tanan/moniple/application"
	"github.com/tanan/moniple/domain/repository"
	"github.com/tanan/moniple/infrastracture/persistence/datastore"
	"github.com/tanan/moniple/interfaces/api/handler"
)

type Registry struct {
	useCase         usecase.ScheduleUseCase
	repository      repository.ScheduleRepository
	ScheduleHandler handler.ScheduleHandler
	MonitorHandler  handler.MonitorHandler
}

func NewRegistry(sqlHandler *datastore.SQLHandler) *Registry {
	repository := NewScheduleRepository(sqlHandler)
	usecase := NewScheduleUseCase(repository)
	return &Registry{
		repository:      repository,
		useCase:         usecase,
		ScheduleHandler: handler.NewScheduleHandler(usecase),
		MonitorHandler:  handler.NewMonitorHandler(),
	}
}

func NewScheduleUseCase(repository repository.ScheduleRepository) usecase.ScheduleUseCase {
	return usecase.NewScheduleUseCase(repository)
}

func NewScheduleRepository(handler *datastore.SQLHandler) repository.ScheduleRepository {
	return datastore.NewScheduleRepository(handler)
}
