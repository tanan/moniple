package schema

type Schedule struct {
	Id       int64
	Name     string
	URL      string
	Method   string
	Interval int64
	Status   string
}
