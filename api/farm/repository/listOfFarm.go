package repository

import (
	"delos/api/model"

	"github.com/jinzhu/gorm"
)

func (repository) ListOfFarm(db *gorm.DB) (*model.Farms, error) {

	farm := new(model.Farm)
	var listFarms model.Farms

	q := db.Table(farm.TableName())

	err := q.Preload("Ponds").Where("deleted_at is NULL").Find(&listFarms).Error

	if err != nil {
		return nil, err
	}

	return &listFarms, nil
}
