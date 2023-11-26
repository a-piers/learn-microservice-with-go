package database

import (
	"database/sql"
	"fmt"
	"lib/utils"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	Cursor *gorm.DB
	DB     *sql.DB
}

func NewStorage() (*Storage, error) {
	host := os.Getenv("DBHOST")
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("DBPORT")
	sslmode := os.Getenv("DBSSLMODE")

	connection_str := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)

	cursor, err := gorm.Open(postgres.Open(connection_str), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error on init database! %s", utils.ExceptionToString(err))
	}

	db, err := cursor.DB()
	if err != nil {
		return nil, fmt.Errorf("error occurred on initializing database! %s", utils.ExceptionToString(err))
	}

	return &Storage{
		Cursor: cursor,
		DB:     db,
	}, nil
}

func (stg *Storage) AutoMigration(dst ...interface{}) error {
	if err := stg.Cursor.AutoMigrate(dst...); err != nil {
		return err
	}

	return nil
}

func (stg *Storage) GetCursor() *gorm.DB {
	return stg.Cursor
}

func (stg *Storage) Ping() error {
	return stg.DB.Ping()
}

func (stg *Storage) Close() error {
	return stg.DB.Close()
}
