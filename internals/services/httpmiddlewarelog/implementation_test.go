package httpmiddlewarelog_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gndw/starting-golang/internals/constants"
	"github.com/gndw/starting-golang/internals/services/httpmiddlewarelog"
	log_mocks "github.com/gndw/starting-golang/internals/services/log/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewHttpMiddlewareService(t *testing.T) {
	type mocks struct {
		logService *log_mocks.Service
	}
	tests := []struct {
		name    string
		setup   func(m mocks)
		wantErr bool
	}{
		{
			name: "should successfully create new http middleware service",
			setup: func(m mocks) {
				// No expectations needed
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mocks{
				logService: log_mocks.NewService(t),
			}
			if tt.setup != nil {
				tt.setup(m)
			}
			got, err := httpmiddlewarelog.NewHttpMiddlewareService(context.Background(), m.logService)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHttpMiddlewareService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)
		})
	}
}

func TestLogMiddleware(t *testing.T) {
	type mocks struct {
		logService *log_mocks.Service
	}
	tests := []struct {
		name         string
		setup        func(m mocks)
		method       string
		path         string
		handlerFunc  constants.HttpFunction
		wantResponse interface{}
		wantErr      bool
	}{
		{
			name: "should successfully log incoming request and return response",
			setup: func(m mocks) {
				m.logService.EXPECT().Info(mock.Anything, "[incoming-http]", []interface{}{"method", "GET", "path", "/test-path"}).Return()
			},
			method: "GET",
			path:   "/test-path",
			handlerFunc: func(ctx context.Context, w http.ResponseWriter, r *http.Request) (interface{}, error) {
				return "test-response", nil
			},
			wantResponse: "test-response",
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mocks{
				logService: log_mocks.NewService(t),
			}
			if tt.setup != nil {
				tt.setup(m)
			}
			s, _ := httpmiddlewarelog.NewHttpMiddlewareService(context.Background(), m.logService)

			req := httptest.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()

			middleware := s.LogMiddleware(tt.handlerFunc)
			resp, err := middleware(context.Background(), w, req)

			if (err != nil) != tt.wantErr {
				t.Errorf("LogMiddleware() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantResponse, resp)
		})
	}
}
