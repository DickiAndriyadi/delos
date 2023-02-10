package service

import (
	"delos/api/model"
	"delos/api/model/response"
	"delos/api/pond/filter"
	"errors"
	"net/http"
)

func (s *service) ListOfPond(f *filter.FilterListOfPond) (*model.Ponds, *response.ErrorResponse) {

	res, err := s.pondRepository.ListOfPond(s.dbManager.GetDB(), f)
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
