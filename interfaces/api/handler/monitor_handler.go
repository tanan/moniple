package handler

import (
	"github.com/labstack/echo"
	"github.com/tanan/moniple/interfaces/api"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type MonitorHandler interface {
	CheckEndpoint(c echo.Context) error
}

type monitorHandler struct {
}

func NewMonitorHandler() MonitorHandler {
	return &monitorHandler{}
}

func (handler *monitorHandler) CheckEndpoint(c echo.Context) error {
	m := new(api.MonitorRequest)
	if err := c.Bind(&m); err != nil {
		return err
	}
	req, _ := http.NewRequest(m.Method, m.URL, nil)
	client := http.Client{
		Timeout: time.Duration(10) * time.Second,
	}
	// time start
	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	duration := time.Now().Sub(start)
	defer resp.Body.Close()
	res := &api.MonitorResponse{
		Status:    http.StatusOK,
		Content:   handler.getContent(resp.Body),
		TimeTaken: duration.Seconds(),
	}
	return c.JSON(resp.StatusCode, res)
}

func (handler *monitorHandler) getContent(body io.ReadCloser) string {
	b, err := ioutil.ReadAll(body)
	if err == nil {
		return string(b)
	}
	return ""
}
