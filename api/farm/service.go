package farm

import (
	"delos/api/model"
	"delos/api/model/payload"
	"delos/api/model/response"
)

type Service interface {
	AddNewFarm(p payload.AddUpdateFarm) (*model.Farm, *response.ErrorResponse)
	ListOfFarm() (*model.Farms, *response.ErrorResponse)
	DetailOfFarm(id string) (*model.Farm, *response.ErrorResponse)
	UpdateFarm(p payload.AddUpdateFarm, id string) (*model.Farm, *response.ErrorResponse)
	DeleteFarm(id string) (string, *response.ErrorResponse)
}
