package database

type SqlHandler interface {
	FindAll(out interface{}) error
	Find(out interface{}, where ...interface{}) error
	Create(value interface{}) error
	NewRecord(value interface{}) bool
	Save(value interface{}) error
}
