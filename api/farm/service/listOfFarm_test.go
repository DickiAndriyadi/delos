package service

import (
	"delos/api/model"
	"delos/api/model/response"
	"errors"
	"net/http"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func Test_service_ListOfFarm(t *testing.T) {
	var (
		db = gorm.DB{}
	)
	tests := []struct {
		name       string
		expectFunc func(t *TestingService)
		want       *model.Farms
		want1      *response.ErrorResponse
	}{
		// Add test cases here ...
		{
			name: "Success Case - ListOfFarm",
			expectFunc: func(t *TestingService) {
				t.farmRepository.On("ListOfFarm", &db).Return(&model.Farms{
					{
						ID: 1,
					},
				}, nil)

			},
			want: &model.Farms{
				{
					ID: 1,
				},
			},
			want1: nil,
		},

		{
			name: "Failed Case - Data not found",
			expectFunc: func(t *TestingService) {
				t.farmRepository.On("ListOfFarm", &db).Return(nil, errors.New("data not found!"))
			},
			want1: &response.ErrorResponse{
				ErrCode: http.StatusInternalServerError,
				Message: errors.New("data not found!"),
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

			got, got1 := testing.Service.ListOfFarm()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
