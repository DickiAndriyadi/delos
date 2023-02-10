package pond

import (
	"delos/api/model"
	"delos/api/model/payload"
	"delos/api/pond/filter"

	"github.com/jinzhu/gorm"
)

type Repository interface {
	AddNewPond(db *gorm.DB, store *model.Pond) (*model.Pond, error)
	ListOfPond(db *gorm.DB, f *filter.FilterListOfPond) (*model.Ponds, error)
	DetailOfPond(db *gorm.DB, payload map[string]interface{}) (*model.Pond, error)
	UpdatePond(db *gorm.DB, store *payload.AddUpdatePond, id string) (*model.Pond, error)
	DeletePond(db *gorm.DB, deletePondMap map[string]interface{}) error
}
