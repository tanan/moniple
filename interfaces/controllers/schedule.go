package controllers

import "github.com/tanan/moniple/interfaces/database"

type ScheduleController struct {

}

func NewScheduleController(sqlHandler database.SqlHandler) *ScheduleController {
	return &ScheduleController{}
}