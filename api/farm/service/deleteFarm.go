package service

import (
	"delos/api/model/response"
	"errors"
	"net/http"

	"github.com/jinzhu/gorm"
)

func (s *service) DeleteFarm(id string) (string, *response.ErrorResponse) {

	detailMap := map[string]interface{}{
		"id": id,
	}

	_, err := s.farmRepository.DetailOfFarm(s.dbManager.GetDB(), detailMap)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return "", &response.ErrorResponse{
				ErrCode: http.StatusBadRequest,
				Message: errors.New("data not found!"),
			}
		}
		return "", &response.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			Message: err,
		}

	}

	deletePondMap := map[string]interface{}{
		"farm_id": id,
	}

	_ = s.pondRepository.DeletePond(s.dbManager.GetDB(), deletePondMap)

	errDelete := s.farmRepository.DeleteFarm(s.dbManager.GetDB(), id)
	if errDelete != nil {
		return "", &response.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			Message: errDelete,
		}
	}

	return "success delete data!", nil
}
