package service

import (
	"delos/api/model"
	"delos/api/model/response"
	"errors"
	"net/http"

	"github.com/jinzhu/gorm"
)

func (s *service) DetailOfFarm(id string) (*model.Farm, *response.ErrorResponse) {

	detailMap := map[string]interface{}{
		"id": id,
	}

	res, err := s.farmRepository.DetailOfFarm(s.dbManager.GetDB(), detailMap)

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

	return res, nil
}
