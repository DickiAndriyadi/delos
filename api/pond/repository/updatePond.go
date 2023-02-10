package repository

import (
	"delos/api/model"
	"delos/api/model/payload"

	"github.com/jinzhu/gorm"
)

func (repository) UpdatePond(db *gorm.DB, store *payload.AddUpdatePond, id string) (*model.Pond, error) {

	var pond model.Pond

	err := db.First(&pond, id).Model(&pond).Updates(store).Error
	if err != nil {
		return nil, err
	}

	return &pond, nil
}
