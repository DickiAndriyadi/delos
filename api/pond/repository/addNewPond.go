package repository

import (
	"delos/api/model"

	"github.com/jinzhu/gorm"
)

func (repository) AddNewPond(db *gorm.DB, store *model.Pond) (*model.Pond, error) {

	var (
		pond model.Pond
	)

	err := db.Table(pond.TableName()).Create(&store).Last(&pond).Error
	if err != nil {
		return nil, err
	}

	return &pond, nil
}
