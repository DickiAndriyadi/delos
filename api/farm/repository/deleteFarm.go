package repository

import (
	"delos/api/model"

	"github.com/jinzhu/gorm"
)

func (repository) DeleteFarm(db *gorm.DB, id string) error {

	var (
		farm model.Farm
	)

	err := db.Where("id = ?", id).Delete(&farm).Error
	if err != nil {
		return err
	}

	return nil
}
