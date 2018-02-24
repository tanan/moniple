package domain

type ScheduleId int64

type Schedule struct {
	Id       ScheduleId
	Name     string
	URL      string
	Interval int64
	Active   bool
}
