package controllers

import (
	"github.com/tanan/moniple/domain"
	"github.com/tanan/moniple/interfaces/database"
	"github.com/tanan/moniple/interfaces/model"
	"github.com/tanan/moniple/usecase/interactor"
)

type ScheduleController struct {
	Interactor interactor.ScheduleInteractor
}

func NewScheduleController(sqlHandler database.SqlHandler) *ScheduleController {
	return &ScheduleController{
		Interactor: interactor.ScheduleInteractor{
			Repository: database.ScheduleRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (c ScheduleController) Store(context Context) {
	err := c.Interactor.Store(domain.Schedule{})
	if err != nil {
		context.JSON(500, model.ErrorResponse{
			Status:  500,
			Message: "Failed to store a schedule.",
		})
	}
	context.JSON(200, model.ScheduleResponse{
		Status: 200,
	})
}
