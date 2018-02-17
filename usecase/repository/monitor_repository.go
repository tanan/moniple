package repository

import "github.com/tanan/moniple/domain"

type MonitorRepository interface {
	CheckEndpointStatus(url string) (domain.MonitorResult, error)
}
