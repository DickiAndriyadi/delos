package service

import (
	"delos/api/model"
	"delos/api/model/payload"
	"delos/api/model/response"
	"errors"
	"net/http"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_service_UpdatePond(t *testing.T) {
	type args struct {
		id      string
		payload payload.AddUpdatePond
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
			name: "Success Case - UpdatePond",
			args: args{
				id: "1",
				payload: payload.AddUpdatePond{
					Title: "pond 1",
				},
			},
			expectFunc: func(t *TestingService) {
				t.pondRepository.On("DetailOfPond", &db, mock.Anything).Return(&model.Pond{
					ID: 1,
				}, nil)
				t.pondRepository.On("UpdatePond", &db, mock.Anything, mock.Anything).Return(&model.Pond{ID: 1}, nil)

			},
			want:  &model.Pond{ID: 1},
			want1: nil,
		},

		{
			name: "Failed Case - Internal server error",
			args: args{
				id: "1",
				payload: payload.AddUpdatePond{
					Title: "pond 1",
				},
			},
			expectFunc: func(t *TestingService) {
				t.pondRepository.On("DetailOfPond", &db, mock.Anything).Return(&model.Pond{
					ID: 1,
				}, nil)
				t.pondRepository.On("UpdatePond", &db, mock.Anything, mock.Anything).Return(nil, errors.New("internal server error"))
			},
			want: nil,
			want1: &response.ErrorResponse{
				ErrCode: http.StatusInternalServerError,
				Message: errors.New("internal server error"),
			},
		},

		{
			name: "Failed Case - Data not found",
			args: args{
				id: "1",
				payload: payload.AddUpdatePond{
					Title: "pond 1",
				},
			},
			expectFunc: func(t *TestingService) {
				t.pondRepository.On("DetailOfPond", &db, mock.Anything).Return(nil, errors.New("data not found"))
			},
			want1: &response.ErrorResponse{
				ErrCode: http.StatusInternalServerError,
				Message: errors.New("data not found"),
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

			got, got1 := testing.Service.UpdatePond(tt.args.payload, tt.args.id)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
