package rest

import (
	"github.com/gin-gonic/gin"
	//"github.com/bamzi/jobrunner"
	"github.com/tanan/moniple/interfaces/model"
	"github.com/tanan/moniple/interfaces/controllers"
	"github.com/tanan/moniple/infrastracture/database"
	"github.com/tanan/moniple/config"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	sqlHandler := database.NewSqlHandler(config.GetDBConnInfo())
	scheduleController := controllers.NewScheduleController(sqlHandler)
	monitorController := controllers.NewMonitorController()

	router.GET("/system", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	//router.GET("/jobrunner/status", func(c *gin.Context) {
	//	c.JSON(200, jobrunner.StatusJson())
	//})

	router.GET("/endpoint", func(c *gin.Context) {
		url := c.Query("url")
		if url == "" {
			c.JSON(500, gin.H{
				"message": "url parameter is set",
			})
			return
		}
		monitorController.Get(model.MonitorRequest{
			URL: url,
			Method: "GET",
		}, c)
	})

	Router = router
}
