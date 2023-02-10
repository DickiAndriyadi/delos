package repository

import (
	"delos/api/model"

	"github.com/jinzhu/gorm"
)

func (repository) DetailOfPond(db *gorm.DB, payload map[string]interface{}) (*model.Pond, error) {

	pond := new(model.Pond)

	q := db.Model(pond)

	err := q.Where(payload).Find(&pond).Error

	if err != nil {
		return nil, err
	}

	return pond, nil
}
