package api

type MonitorRequest struct {
	URL    string
	Method string
}

type ScheduleRequest struct {
	ID       string
	Name     string
	URL      string
	Method   string
	Interval int64
	Status   string
}
