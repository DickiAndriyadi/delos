package db

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/sirupsen/logrus"
)

type DatabaseManager interface {
	GetDB() *gorm.DB
	Initialize(dsn string, connection string) error
}

func NewDatabaseManager() DatabaseManager {
	return &databaseManager{}
}

type databaseManager struct {
	db *gorm.DB
}

const maxTriedOpenDb = 3

func createDBInstance(dsn string, connection string, tried int) (*gorm.DB, error) {

	db, err := gorm.Open(connection, dsn)

	if err != nil {
		logrus.Error(err.Error())
		panic("failed to connect database")
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour * 2)
	db.SetLogger(&GormLogger{})
	db.LogMode(true)

	if maxTriedOpenDb == tried {
		logrus.Error("Failed to open database")
		return nil, err
	}

	if err := db.DB().Ping(); err != nil {
		db.Close()
		return createDBInstance(dsn, connection, tried+1)
	}

	return db, nil
}

func (dbManager *databaseManager) Initialize(dsn string, connection string) error {
	var err error

	if dbManager.db, err = createDBInstance(dsn, connection, 0); err != nil {
		return err
	}

	return nil
}

func (dbManager *databaseManager) GetDB() *gorm.DB {
	if dbManager.db == nil {
		return nil
	}
	return dbManager.db
}

type GormLogger struct {
}

func (g *GormLogger) Print(v ...interface{}) {
	switch v[0] {
	case "sql":
		logrus.Infof("[db-log][%v][%v][rows_returned:%v]{query:\"%v\", values:\"%v\", src:\"%v\"}\n",
			time.Now().Format("2006-01-02T15:04:05.000000Z"),
			v[2],
			v[5],
			v[3],
			v[4],
			v[1],
		)
	case "log":
		logrus.Infof("[db-log][%v][%v]\n",
			time.Now().Format("2006-01-02T15:04:05.000000Z"),
			v[2],
		)
	}
}
