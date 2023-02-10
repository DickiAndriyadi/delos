package repository

import (
	"delos/api/model"
	"delos/api/pond/filter"

	"github.com/jinzhu/gorm"
)

func (repository) ListOfPond(db *gorm.DB, f *filter.FilterListOfPond) (*model.Ponds, error) {

	pond := new(model.Pond)
	var listPonds model.Ponds

	q := db.Table(pond.TableName())

	if f != nil {
		if f.FarmID != "" {
			q = q.Where("farm_id = ?", f.FarmID)
		}
	}

	err := q.Where("deleted_at is NULL").Find(&listPonds).Error

	if err != nil {
		return nil, err
	}

	return &listPonds, nil
}
