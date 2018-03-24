package api

type MonitorResponse struct {
	Status    int
	TimeTaken float64
	Content   string
}

type ScheduleResponse struct {
	ID       int
	Name     string
	URL      string
	Method   string
	Interval int64
	Status   string
}

type ErrorResponse struct {
	Status  int
	Message string
}
