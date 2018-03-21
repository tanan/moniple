package datastore

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type SQLHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler(connInfo string) *SQLHandler {
	handler := &SQLHandler{}
	err := handler.open(connInfo)
	if err != nil {
		panic(err)
	}
	return handler
}

func (handler *SQLHandler) open(connInfo string) error {
	conn, err := gorm.Open("mysql", connInfo)
	if err != nil {
		return err
	}
	handler.Conn = conn
	return nil
}

func (handler *SQLHandler) FindAll(out interface{}) error {
	handler.Conn.Find(out)
	return nil
}

func (handler *SQLHandler) Find(out interface{}, where ...interface{}) error {
	if len(where) > 0 {
		return handler.Conn.Find(out, where...).Error
	} else {
		return handler.Conn.Find(out).Error
	}
}

func (handler *SQLHandler) Create(value interface{}) error {
	if err := handler.Conn.Create(value).Error; err != nil {
		return err
	}
	return nil
}

func (handler *SQLHandler) Save(value interface{}) error {
	if err := handler.Conn.Save(value).Error; err != nil {
		return err
	}
	return nil
}

func (handler *SQLHandler) NewRecord(value interface{}) bool {
	return handler.Conn.NewRecord(value)
}
