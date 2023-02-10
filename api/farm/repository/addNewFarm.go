package repository

import (
	"delos/api/model"

	"github.com/jinzhu/gorm"
)

func (repository) AddNewFarm(db *gorm.DB, store *model.Farm) (*model.Farm, error) {

	var (
		farm model.Farm
	)

	err := db.Table(farm.TableName()).Create(&store).Last(&farm).Error
	if err != nil {
		return nil, err
	}

	return &farm, nil
}
