package farm

import (
	"delos/api/model"
	"delos/api/model/payload"

	"github.com/jinzhu/gorm"
)

type Repository interface {
	AddNewFarm(db *gorm.DB, store *model.Farm) (*model.Farm, error)
	ListOfFarm(db *gorm.DB) (*model.Farms, error)
	DetailOfFarm(db *gorm.DB, payload map[string]interface{}) (*model.Farm, error)
	UpdateFarm(db *gorm.DB, store *payload.AddUpdateFarm, id string) (*model.Farm, error)
	DeleteFarm(db *gorm.DB, id string) error
}
