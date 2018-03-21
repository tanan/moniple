package repository

import "github.com/tanan/moniple/domain/model"

type MonitorRepository interface {
	//CheckEndpointStatus(url string) (model.MonitorResult, error)
	FindEndpointByID(id model.MonitorId) (model.Monitor, error)
}
