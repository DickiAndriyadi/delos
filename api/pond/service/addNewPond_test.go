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

func Test_service_AddNewPond(t *testing.T) {
	type args struct {
		payload payload.AddUpdatePond
	}
	var (
		db        = gorm.DB{}
		valuePond = model.Pond{
			ID:          1,
			Title:       "pond 1",
			Description: "pond 1 desc",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil,
		}
	)
	tests := []struct {
		name       string
		args       args
		expectFunc func(t *TestingService)
		want       *model.Pond
		want1      *response.ErrorResponse
	}{
		// Add test cases here ...
		{
			name: "Success Case - AddNewPond",
			args: args{
				payload: payload.AddUpdatePond{
					Title:  "pond 1",
					FarmID: 1,
					Req:    "Add",
				},
			},
			expectFunc: func(t *TestingService) {
				t.farmRepository.On("DetailOfFarm", &db, mock.Anything).Return(&model.Farm{ID: 1}, nil)
				t.pondRepository.On("DetailOfPond", &db, mock.Anything).Return(nil, errors.New("data not found"))
				t.pondRepository.On("AddNewPond", &db, mock.Anything).Return(&valuePond, nil)
			},
			want:  &valuePond,
			want1: nil,
		},

		{
			name: "Failed Case - Title already used",
			args: args{
				payload: payload.AddUpdatePond{
					Title:  "pond 1",
					FarmID: 1,
					Req:    "Add",
				},
			},
			expectFunc: func(t *TestingService) {
				t.farmRepository.On("DetailOfFarm", &db, mock.Anything).Return(&model.Farm{ID: 1}, nil)
				t.pondRepository.On("DetailOfPond", &db, mock.Anything).Return(&valuePond, nil)
			},
			want: nil,
			want1: &response.ErrorResponse{
				ErrCode: http.StatusBadRequest,
				Message: errors.New("Pond with title pond 1 is already used"),
			},
		},

		{
			name: "Failed Case - Internal Server Error",
			args: args{
				payload: payload.AddUpdatePond{
					Title:  "pond 1",
					FarmID: 1,
					Req:    "Add",
				},
			},
			expectFunc: func(t *TestingService) {
				t.farmRepository.On("DetailOfFarm", &db, mock.Anything).Return(&model.Farm{ID: 1}, nil)
				t.pondRepository.On("DetailOfPond", &db, mock.Anything).Return(nil, errors.New("data not found!"))
				t.pondRepository.On("AddNewPond", &db, mock.Anything).Return(nil, errors.New("internal server error"))

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

			got, got1 := testing.Service.AddNewPond(tt.args.payload)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
