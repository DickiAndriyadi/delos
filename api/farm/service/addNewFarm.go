package service

import (
	"delos/api/model"
	"delos/api/model/payload"
	"delos/api/model/response"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
)

func (s *service) AddNewFarm(p payload.AddUpdateFarm) (*model.Farm, *response.ErrorResponse) {

	detailMap := map[string]interface{}{
		"title": p.Title,
	}

	_, err := s.farmRepository.DetailOfFarm(s.dbManager.GetDB(), detailMap)

	if err == nil {
		if !gorm.IsRecordNotFoundError(err) {
			return nil, &response.ErrorResponse{
				ErrCode: http.StatusBadRequest,
				Message: errors.New(fmt.Sprintf("Farm with title %s is already used", p.Title)),
			}
		}
	}

	var (
		farm model.Farm
	)
	farm.Title = p.Title
	farm.Description = p.Description
	farm.CreatedAt = time.Now()
	farm.UpdatedAt = time.Now()

	data, errCreate := s.farmRepository.AddNewFarm(s.dbManager.GetDB(), &farm)
	if errCreate != nil {
		return nil, &response.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			Message: errCreate,
		}
	}

	return data, nil
}
