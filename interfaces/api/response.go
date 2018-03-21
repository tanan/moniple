package api

type MonitorResponse struct {
	Status    int
	TimeTaken int
	Content   string
}

type ScheduleResponse struct {
	Status int
}

type ErrorResponse struct {
	Status  int
	Message string
}
