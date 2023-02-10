package repository

import (
	"delos/api/model"

	"github.com/jinzhu/gorm"
)

func (repository) DetailOfFarm(db *gorm.DB, payload map[string]interface{}) (*model.Farm, error) {

	farm := new(model.Farm)

	q := db.Model(farm)

	err := q.Preload("Ponds").Where(payload).Find(&farm).Error

	if err != nil {
		return nil, err
	}

	return farm, nil
}
