package service

import (
	"delos/api/model/response"
	"errors"
	"net/http"

	"github.com/jinzhu/gorm"
)

func (s *service) DeletePond(id string) (string, *response.ErrorResponse) {

	detailMap := map[string]interface{}{
		"id": id,
	}

	_, err := s.pondRepository.DetailOfPond(s.dbManager.GetDB(), detailMap)

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
		"id": id,
	}

	errDelete := s.pondRepository.DeletePond(s.dbManager.GetDB(), deletePondMap)
	if errDelete != nil {
		return "", &response.ErrorResponse{
			ErrCode: http.StatusInternalServerError,
			Message: errDelete,
		}
	}

	return "success delete data!", nil
}
