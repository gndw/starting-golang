package test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gndw/starting-golang/internals/models"
	httpserver_mocks "github.com/gndw/starting-golang/internals/services/httpserver/mocks"
	test_usecase_mocks "github.com/gndw/starting-golang/internals/usecase/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestImplementation_Test(t *testing.T) {
	type mocks struct {
		usecase *test_usecase_mocks.Usecase
	}
	type args struct {
		ctx     context.Context
		request models.TestRequest
		isNilBody bool
	}
	tests := []struct {
		name    string
		setup   func(m mocks)
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "should return success when usecase returns success",
			args: args{
				ctx: context.Background(),
				request: models.TestRequest{
					UserID: 100,
				},
			},
			setup: func(m mocks) {
				m.usecase.EXPECT().Test(mock.Anything, models.TestRequest{UserID: 100}).Return(models.TestResponse{FullName: "John Doe"}, nil)
			},
			want:    models.TestResponse{FullName: "John Doe"},
			wantErr: false,
		},
		{
			name: "should return error when usecase returns error",
			args: args{
				ctx: context.Background(),
				request: models.TestRequest{
					UserID: 100,
				},
			},
			setup: func(m mocks) {
				m.usecase.EXPECT().Test(mock.Anything, models.TestRequest{UserID: 100}).Return(models.TestResponse{}, errors.New("usecase error"))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return error when json decoding fails",
			args: args{
				ctx: context.Background(),
				isNilBody: true,
			},
			setup:   func(m mocks) {},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mocks{
				usecase: test_usecase_mocks.NewUsecase(t),
			}
			if tt.setup != nil {
				tt.setup(m)
			}
			h := &Implementation{
				testUsecase: m.usecase,
			}

			var r *http.Request
			if tt.args.isNilBody {
				r = httptest.NewRequest(http.MethodPost, "/test", nil)
			} else {
				body, _ := json.Marshal(tt.args.request)
				r = httptest.NewRequest(http.MethodPost, "/test", bytes.NewBuffer(body))
			}
			rw := httptest.NewRecorder()

			got, err := h.Test(tt.args.ctx, rw, r)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestNewHandler(t *testing.T) {
	t.Run("should successfully create new handler", func(t *testing.T) {
		httpService := httpserver_mocks.NewService(t)
		usecase := test_usecase_mocks.NewUsecase(t)
		
		httpService.EXPECT().RegisterEndpoint(mock.Anything, http.MethodPost, "/test", mock.Anything).Return(nil)
		
		h, err := NewHandler(context.Background(), httpService, usecase)
		assert.NoError(t, err)
		assert.NotNil(t, h)
		impl := h.(*Implementation)
		assert.Equal(t, usecase, impl.testUsecase)
	})
}
