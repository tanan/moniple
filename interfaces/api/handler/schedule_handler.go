package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/tanan/moniple/application"
	"github.com/tanan/moniple/domain/model"
	"github.com/tanan/moniple/interfaces/api"
	"net/http"
	"strconv"
)

type ScheduleHandler interface {
	SaveSchedule(c echo.Context) error
	GetSchedule(c echo.Context) error
	ListSchedule(c echo.Context) error
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

func (handler *scheduleHandler) GetSchedule(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "schedule id should be numeric value.")
	}
	schedule, err := handler.scheduleUseCase.FindScheduleById(model.ScheduleID(id))
	if err != nil || schedule.Name == "" {
		return c.String(http.StatusNoContent, "schedule is not found")
	}

	response := &api.ScheduleResponse{
		ID:       int(schedule.ID),
		Name:     schedule.Name,
		URL:      schedule.URL,
		Method:   schedule.Method,
		Interval: schedule.Interval,
		Status:   schedule.GetActive(),
	}
	return c.JSON(http.StatusOK, response)
}

func (handler *scheduleHandler) ListSchedule(c echo.Context) error {
	schedules, err := handler.scheduleUseCase.FindAll()
	if err != nil || len(schedules) == 0 {
		return c.String(http.StatusNoContent, "schedule is not found.")
	}

	var response []*api.ScheduleResponse

	for _, schedule := range schedules {
		r := &api.ScheduleResponse{
			ID:       int(schedule.ID),
			Name:     schedule.Name,
			URL:      schedule.URL,
			Method:   schedule.Method,
			Interval: schedule.Interval,
			Status:   schedule.GetActive(),
		}
		response = append(response, r)
	}
	return c.JSON(http.StatusOK, response)
}
