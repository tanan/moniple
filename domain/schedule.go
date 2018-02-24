package domain

type ScheduleId int64

type Schedule struct {
	Id       ScheduleId
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
