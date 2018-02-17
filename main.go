package main

import (
	"github.com/tanan/moniple/infrastracture/rest"
	"github.com/tanan/moniple/config"
	"strconv"
)

func main() {
	rest.Router.Run(":" + strconv.Itoa(config.GetAPIPort()))
}