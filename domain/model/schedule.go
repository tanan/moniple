package model

type ScheduleID int64

type Schedule struct {
	ID       ScheduleID
	Name     string
	URL      string
	Method   string
	Interval int64
	Active   bool
}

func (s *Schedule) SetActive(status string) {
	if status == "active" {
		s.Active = true
		return
	}
	s.Active = false
}
