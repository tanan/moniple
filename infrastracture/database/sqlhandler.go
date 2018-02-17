package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler(connInfo string) *SqlHandler {
	handler := &SqlHandler{}
	err := handler.open(connInfo)
	if err != nil {
		panic(err)
	}
	return handler
}

func (handler *SqlHandler) open(connInfo string) error {
	conn, err := gorm.Open("mysql", connInfo)
	if err != nil {
		return err
	}
	handler.Conn = conn
	return nil
}

func (handler *SqlHandler) FindAll(out interface{}) error {
	handler.Conn.Find(out)
	return nil
}

func (handler *SqlHandler) Find(out interface{}, where ...interface{}) error {
	if len(where) > 0 {
		handler.Conn.Find(out, where...)
	} else {
		handler.Conn.Find(out)
	}
	return nil
}
