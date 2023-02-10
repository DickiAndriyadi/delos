package service

import (
	"delos/api/model"
	"delos/api/model/payload"
	"delos/api/model/response"
	"errors"
	"net/http"

	"github.com/jinzhu/gorm"
)

func (s *service) UpdateFarm(p payload.AddUpdateFarm, id string) (*model.Farm, *response.ErrorResponse) {

	detailMap := map[string]interface{}{
		"id": id,
	}

	_, err := s.farmRepository.DetailOfFarm(s.dbManager.GetDB(), detailMap)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, &response.ErrorResponse{
				ErrCode: http.StatusBadRequest,
				Message: errors.New("data not found!"),
			}
		}

		return nil, &response.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			Message: err,
		}
	}

	data, errUpdate := s.farmRepository.UpdateFarm(s.dbManager.GetDB(), &p, id)
	if errUpdate != nil {

		return nil, &response.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			Message: errUpdate,
		}
	}

	return data, nil
}
