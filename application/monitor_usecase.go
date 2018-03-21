package usecase

import (
	"github.com/tanan/moniple/domain/model"
	"github.com/tanan/moniple/domain/repository"
)

type MonitorUseCase interface {
	CheckEndpoint(url string) (model.MonitorResult, error)
}

type monitorUseCase struct {
	Repository repository.MonitorRepository
}

func (i monitorUseCase) CheckEndpoint(id int) (model.MonitorResult, error) {
	//monitor, err := i.Repository.FindEndpointByID(model.MonitorId(id))
	_, err := i.Repository.FindEndpointByID(model.MonitorId(id))
	if err != nil {
		return model.MonitorResult{}, err
	}
	// TODO Monitorリクエスト実装
	return model.MonitorResult{}, err
}
