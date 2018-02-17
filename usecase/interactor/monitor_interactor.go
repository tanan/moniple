package interactor

import (
	"github.com/tanan/moniple/usecase/repository"
	"github.com/tanan/moniple/domain"
)

type MonitorInteractor struct {
	Repository repository.MonitorRepository
}

func (i MonitorInteractor) CheckEndpoint(url string) (domain.MonitorResult, error) {
	return i.Repository.CheckEndpointStatus(url)
}