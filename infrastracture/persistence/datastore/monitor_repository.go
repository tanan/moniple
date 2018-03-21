package datastore

import "github.com/tanan/moniple/domain/model"

type MonitorRepository struct {
	*SQLHandler
}

func (r *MonitorRepository) FindEndpointByID(id model.MonitorId) (model.Monitor, error) {
	var m model.Monitor
	err := r.Conn.Find(&m, "id = ?", id).Error
	if err != nil {
		return model.Monitor{}, err
	}
	return m, nil
}
