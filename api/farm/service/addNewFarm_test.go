package service

import (
	"delos/api/model"
	"delos/api/model/payload"
	"delos/api/model/response"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_service_AddNewFarm(t *testing.T) {
	type args struct {
		payload payload.AddUpdateFarm
	}
	var (
		db        = gorm.DB{}
		valueFarm = model.Farm{
			ID:          1,
			Title:       "farm 1",
			Description: "farm 1 desc",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil,
		}
	)
	tests := []struct {
		name       string
		args       args
		expectFunc func(t *TestingService)
		want       *model.Farm
		want1      *response.ErrorResponse
	}{
		// Add test cases here ...
		{
			name: "Success Case - AddNewFarm",
			args: args{
				payload: payload.AddUpdateFarm{
					Title: "farm 1",
				},
			},
			expectFunc: func(t *TestingService) {
				t.farmRepository.On("DetailOfFarm", &db, mock.Anything).Return(nil, errors.New("data not found"))
				t.farmRepository.On("AddNewFarm", &db, mock.Anything).Return(&valueFarm, nil)
			},
			want:  &valueFarm,
			want1: nil,
		},

		{
			name: "Failed Case - Title already used",
			args: args{
				payload: payload.AddUpdateFarm{
					Title: "farm 1",
				},
			},
			expectFunc: func(t *TestingService) {
				t.farmRepository.On("DetailOfFarm", &db, mock.Anything).Return(&valueFarm, nil)
			},
			want: nil,
			want1: &response.ErrorResponse{
				ErrCode: http.StatusBadRequest,
				Message: errors.New("Farm with title farm 1 is already used"),
			},
		},

		{
			name: "Failed Case - Internal Server Error",
			args: args{
				payload: payload.AddUpdateFarm{
					Title: "farm 1",
				},
			},
			expectFunc: func(t *TestingService) {
				t.farmRepository.On("DetailOfFarm", &db, mock.Anything).Return(nil, errors.New("data not found!"))
				t.farmRepository.On("AddNewFarm", &db, mock.Anything).Return(nil, errors.New("internal server error"))

			},
			want: nil,
			want1: &response.ErrorResponse{
				ErrCode: http.StatusInternalServerError,
				Message: errors.New("internal server error"),
			},
		},

		// End test cases ...
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testing := &TestingService{}
			testing.Initialize(t)

			if tt.expectFunc != nil {
				tt.expectFunc(testing)
			}

			got, got1 := testing.Service.AddNewFarm(tt.args.payload)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
