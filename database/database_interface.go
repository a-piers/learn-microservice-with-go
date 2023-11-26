package database

import "gorm.io/gorm"

type IDatabase interface {
	AutoMigration(...interface{}) error
	GetCursor() *gorm.DB
	Close() error
	Ping() error
}

type CrudInterface interface {
	Get(string) (any, error)
	GetList() ([]any, error)
	Insert(any) error
	Update(any, interface{}) (any, error)
	Delete(string) (any, error)
}
