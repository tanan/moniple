package controllers

import (
	"github.com/tanan/moniple/usecase/interactor"
	"github.com/tanan/moniple/interfaces/model"
	"github.com/tanan/moniple/interfaces/requester"
)

type MonitorController struct {
	Interactor interactor.MonitorInteractor
}

func NewMonitorController() MonitorController {
	return MonitorController{
		Interactor: interactor.MonitorInteractor{
			Repository: requester.MontiorRequester{},
		},
	}
}

func (c MonitorController) Get(req model.MonitorRequest, context Context) {
	resp, err := c.Interactor.CheckEndpoint(req.URL)
	if err != nil {
		context.JSON(500, model.MonitorResponse{
			Status: 500,
			TimeTaken: 0,
			Content: "",
		})
	}
	context.JSON(200, model.MonitorResponse{
		Status: resp.Status,
		TimeTaken: resp.TimeTaken,
		Content: resp.Content,
	})

}