package handler

import (
	"github.com/labstack/echo"
	"github.com/tanan/moniple/application"
	"github.com/tanan/moniple/domain/model"
	"github.com/tanan/moniple/interfaces/api"
	"strconv"
)

type ScheduleHandler interface {
	SaveSchedule(c echo.Context) error
}

type scheduleHandler struct {
	scheduleUseCase usecase.ScheduleUseCase
}

func NewScheduleHandler(usecase usecase.ScheduleUseCase) ScheduleHandler {
	return &scheduleHandler{
		scheduleUseCase: usecase,
	}
}

func (handler *scheduleHandler) SaveSchedule(c echo.Context) error {
	request := new(api.ScheduleRequest)
	if err := c.Bind(&request); err != nil {
		return err
	}
	id, _ := strconv.Atoi(request.ID)
	schedule := model.Schedule{
		ID:       model.ScheduleID(id),
		Name:     request.Name,
		URL:      request.URL,
		Method:   request.Method,
		Interval: request.Interval,
	}
	schedule.SetActive(request.Status)
	handler.scheduleUseCase.Store(schedule)
	return nil
}
