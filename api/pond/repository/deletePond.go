package repository

import (
	"delos/api/model"

	"github.com/jinzhu/gorm"
)

func (repository) DeletePond(db *gorm.DB, deletePondMap map[string]interface{}) error {

	var (
		pond model.Pond
	)

	err := db.Where(deletePondMap).Delete(&pond).Error
	if err != nil {
		return err
	}

	return nil
}
