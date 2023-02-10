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

func (s *service) AddNewPond(p payload.AddUpdatePond) (*model.Pond, *response.ErrorResponse) {

	detailFarmMap := map[string]interface{}{
		"id": p.FarmID,
	}

	_, errFarm := s.farmRepository.DetailOfFarm(s.dbManager.GetDB(), detailFarmMap)

	if errFarm != nil {
		if gorm.IsRecordNotFoundError(errFarm) {
			return nil, &response.ErrorResponse{
				ErrCode: http.StatusBadRequest,
				Message: errors.New(fmt.Sprintf("Farm with ID %d is not found", p.FarmID)),
			}
		}
	}

	detailMap := map[string]interface{}{
		"title": p.Title,
	}

	_, err := s.pondRepository.DetailOfPond(s.dbManager.GetDB(), detailMap)

	if err == nil {
		if !gorm.IsRecordNotFoundError(err) {
			return nil, &response.ErrorResponse{
				ErrCode: http.StatusBadRequest,
				Message: errors.New(fmt.Sprintf("Pond with title %s is already used", p.Title)),
			}
		}
	}

	var (
		pond model.Pond
	)

	pond.FarmID = p.FarmID
	pond.Title = p.Title
	pond.Description = p.Description
	pond.CreatedAt = time.Now()
	pond.UpdatedAt = time.Now()

	data, errCreate := s.pondRepository.AddNewPond(s.dbManager.GetDB(), &pond)
	if errCreate != nil {
		return nil, &response.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			Message: errCreate,
		}
	}

	return data, nil
}
