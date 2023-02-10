package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func initMock() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	gormDB, err := gorm.Open("mysql", db)
	if err != nil {
		panic(err)
	}
	return gormDB, mock
}
