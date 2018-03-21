package rest

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tanan/moniple/config"
	"github.com/tanan/moniple/infrastracture/persistence/datastore"
	"github.com/tanan/moniple/interfaces/registry"
	"net/http"
)

var Router *echo.Echo

func init() {
	router := echo.New()

	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	r := registry.NewRegistry(datastore.NewSqlHandler(config.GetDBConnInfo()))

	//router.GET("/jobrunner/status", func(c *gin.Context) {
	//	c.JSON(200, jobrunner.StatusJson())
	//})

	router.GET("/system", ping)
	router.PUT("/schedule/save", r.Handler.SaveSchedule)

	Router = router
}

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "OK!!")
}
