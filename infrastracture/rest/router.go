package rest

import (
	"github.com/gin-gonic/gin"
	//"github.com/bamzi/jobrunner"
	"github.com/tanan/moniple/config"
	"github.com/tanan/moniple/infrastracture/database"
	"github.com/tanan/moniple/interfaces/controllers"
	"github.com/tanan/moniple/interfaces/model"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	sqlHandler := database.NewSqlHandler(config.GetDBConnInfo())
	scheduleController := controllers.NewScheduleController(sqlHandler)
	monitorController := controllers.NewMonitorController()

	//router.GET("/jobrunner/status", func(c *gin.Context) {
	//	c.JSON(200, jobrunner.StatusJson())
	//})

	router.GET("/system", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	router.GET("/endpoint", func(c *gin.Context) {
		monitorController.Get(model.MonitorRequest{
			URL:    c.Query("url"),
			Method: "GET",
		}, c)
	})

	router.POST("/schedule/:id", func(c *gin.Context) {
		scheduleController.Store(model.ScheduleRequest{
			Id: c.Param("id"),
		}, c)
	})

	Router = router
}
