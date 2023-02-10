package service

import (
	"delos/api/model"
	"delos/api/model/response"
	"delos/api/pond/filter"
	"errors"
	"net/http"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_service_ListOfPond(t *testing.T) {
	var (
		db = gorm.DB{}
	)
	tests := []struct {
		name       string
		expectFunc func(t *TestingService)
		want       *model.Ponds
		want1      *response.ErrorResponse
	}{
		// Add test cases here ...
		{
			name: "Success Case - ListOfPond",
			expectFunc: func(t *TestingService) {
				t.pondRepository.On("ListOfPond", &db, mock.Anything).Return(&model.Ponds{
					{
						ID: 1,
					},
				}, nil)

			},
			want: &model.Ponds{
				{
					ID: 1,
				},
			},
			want1: nil,
		},

		{
			name: "Failed Case - Data not found",
			expectFunc: func(t *TestingService) {
				t.pondRepository.On("ListOfPond", &db, mock.Anything).Return(nil, errors.New("data not found!"))
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

			filter := filter.FilterListOfPond{
				FarmID: "",
			}

			got, got1 := testing.Service.ListOfPond(&filter)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
