package repository

import (
	"delos/api/model"
	"delos/api/model/payload"

	"github.com/jinzhu/gorm"
)

func (repository) UpdateFarm(db *gorm.DB, store *payload.AddUpdateFarm, id string) (*model.Farm, error) {

	var farm model.Farm

	err := db.First(&farm, id).Model(&farm).Updates(store).Error
	if err != nil {
		return nil, err
	}

	return &farm, nil
}
