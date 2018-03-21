package main

import (
	"github.com/tanan/moniple/config"
	"github.com/tanan/moniple/infrastracture/rest"
	"strconv"
)

func main() {
	rest.Router.Logger.Fatal(rest.Router.Start(":" + strconv.Itoa(config.GetAPIPort())))
}
