package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tanan/moniple/domain"
	"github.com/tanan/moniple/interfaces/database"
	"github.com/tanan/moniple/interfaces/model"
	"github.com/tanan/moniple/usecase/interactor"
	"strconv"
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

func (c ScheduleController) Store(req model.ScheduleRequest, context Context) {
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(500, gin.H{
			"message": "Body parameter isn't valid.",
		})
		return
	}

	id, err := strconv.Atoi(req.Id)
	if err != nil {
		context.JSON(500, gin.H{
			"message": "Body parameter isn't valid.",
		})
		return
	}

	var schedule domain.Schedule
	schedule.Id = domain.ScheduleId(id)
	schedule.Name = req.Name
	schedule.URL = req.URL
	schedule.Method = req.Method
	schedule.Interval = req.Interval
	schedule.SetActive(req.Status)

	if err := c.Interactor.Store(schedule); err != nil {
		context.JSON(500, model.ErrorResponse{
			Status:  500,
			Message: "Failed to store a schedule.",
		})
		return
	}
	context.JSON(200, model.ScheduleResponse{
		Status: 200,
	})
}
