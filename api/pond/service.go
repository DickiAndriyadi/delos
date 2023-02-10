package pond

import (
	"delos/api/model"
	"delos/api/model/payload"
	"delos/api/model/response"
	"delos/api/pond/filter"
)

type Service interface {
	AddNewPond(p payload.AddUpdatePond) (*model.Pond, *response.ErrorResponse)
	ListOfPond(f *filter.FilterListOfPond) (*model.Ponds, *response.ErrorResponse)
	DetailOfPond(id string) (*model.Pond, *response.ErrorResponse)
	UpdatePond(p payload.AddUpdatePond, id string) (*model.Pond, *response.ErrorResponse)
	DeletePond(id string) (string, *response.ErrorResponse)
}
