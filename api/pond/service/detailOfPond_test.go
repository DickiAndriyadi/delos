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

func Test_service_DetailOfPond(t *testing.T) {
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
		want       *model.Pond
		want1      *response.ErrorResponse
	}{
		// Add test cases here ...
		{
			name: "Success Case - DetailOfPond",
			args: args{
				id: "1",
			},
			expectFunc: func(t *TestingService) {
				t.pondRepository.On("DetailOfPond", &db, mock.Anything).Return(&model.Pond{
					ID: 1,
				}, nil)

			},
			want: &model.Pond{
				ID: 1,
			},
			want1: nil,
		},

		{
			name: "Failed Case - DetailOfPond",
			args: args{
				id: "2",
			},
			expectFunc: func(t *TestingService) {
				t.pondRepository.On("DetailOfPond", &db, mock.Anything).Return(nil, errors.New("data not found!"))
			},
			want: nil,
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

			got, got1 := testing.Service.DetailOfPond(tt.args.id)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}