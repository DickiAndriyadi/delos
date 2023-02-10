package service

import (
	"delos/api/model"
	"delos/api/model/response"
	"errors"
	"net/http"
)

func (s *service) ListOfFarm() (*model.Farms, *response.ErrorResponse) {

	res, err := s.farmRepository.ListOfFarm(s.dbManager.GetDB())
	if err != nil {
		return nil, &response.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			Message: err,
		}
	}

	if len(*res) < 1 {
		return nil, &response.ErrorResponse{
			ErrCode: http.StatusNotFound,
			Message: errors.New("data not found!"),
		}
	}

	return res, nil
}
