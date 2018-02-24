package model

type MonitorRequest struct {
	URL    string
	Method string
}

type ScheduleRequest struct {
	Id       string
	Name     string
	URL      string
	Method   string
	Interval int64
	Status   string
}
