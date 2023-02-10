package service

import (
	"delos/api/model"
	"delos/api/model/response"
	"errors"
	"net/http"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_service_DeleteFarm(t *testing.T) {
	type args struct {
		id string
	}
	var (
		db = gorm.DB{}
	)
	tests := []struct {
		name       string
		args       args
		expectFunc func(t *TestingService)
		want       string
		want1      *response.ErrorResponse
	}{
		// Add test cases here ...
		{
			name: "Success Case - DeleteFarm",
			args: args{
				id: "1",
			},
			expectFunc: func(t *TestingService) {
				t.farmRepository.On("DetailOfFarm", &db, mock.Anything).Return(&model.Farm{
					ID: 1,
				}, nil)
				t.pondRepository.On("DeletePond", &db, mock.Anything).Return(nil)
				t.farmRepository.On("DeleteFarm", &db, mock.Anything).Return(nil)

			},
			want:  "success delete data!",
			want1: nil,
		},

		{
			name: "Failed Case - Data not found",
			args: args{
				id: "1",
			},
			expectFunc: func(t *TestingService) {
				t.farmRepository.On("DetailOfFarm", &db, mock.Anything).Return(nil, errors.New("data not found"))
			},
			want1: &response.ErrorResponse{
				ErrCode: http.StatusInternalServerError,
				Message: errors.New("data not found"),
			},
		},

		{
			name: "Failed Case - Internal server error",
			args: args{
				id: "1",
			},
			expectFunc: func(t *TestingService) {
				t.farmRepository.On("DetailOfFarm", &db, mock.Anything).Return(&model.Farm{
					ID: 1,
				}, nil)
				t.pondRepository.On("DeletePond", &db, mock.Anything).Return(nil)
				t.farmRepository.On("DeleteFarm", &db, mock.Anything).Return(errors.New("internal server error"))
			},
			want: "",
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

			got, got1 := testing.Service.DeleteFarm(tt.args.id)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
