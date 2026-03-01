package httpserver

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gndw/starting-golang/internals/constants"
	"github.com/gndw/starting-golang/internals/services/env"
	env_mocks "github.com/gndw/starting-golang/internals/services/env/mocks"
	httpmiddlewarelog_mocks "github.com/gndw/starting-golang/internals/services/httpmiddlewarelog/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewHttpServerService(t *testing.T) {
	type mocks struct {
		logMiddleware *httpmiddlewarelog_mocks.Service
		env           *env_mocks.Service
	}
	tests := []struct {
		name    string
		setup   func(m mocks)
		wantErr bool
	}{
		{
			name: "should successfully create new http server service",
			setup: func(m mocks) {
				// No expectations needed
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mocks{
				logMiddleware: httpmiddlewarelog_mocks.NewService(t),
				env:           env_mocks.NewService(t),
			}
			if tt.setup != nil {
				tt.setup(m)
			}
			got, err := NewHttpServerService(context.Background(), m.logMiddleware, m.env)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHttpServerService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)
			assert.NotNil(t, got.handler)
		})
	}
}

func TestRegisterEndpoint(t *testing.T) {
	type mocks struct {
		logMiddleware *httpmiddlewarelog_mocks.Service
		env           *env_mocks.Service
	}
	tests := []struct {
		name           string
		setup          func(m mocks)
		method         string
		path           string
		handlerFunc    constants.HttpFunction
		requestMethod  string
		requestPath    string
		wantStatusCode int
		wantBody       HttpResponse
	}{
		{
			name: "should successfully register and handle endpoint with success response",
			setup: func(m mocks) {
				m.logMiddleware.EXPECT().LogMiddleware(mock.Anything).RunAndReturn(func(f constants.HttpFunction) constants.HttpFunction {
					return f
				})
			},
			method: "GET",
			path:   "/test",
			handlerFunc: func(ctx context.Context, w http.ResponseWriter, r *http.Request) (interface{}, error) {
				return map[string]interface{}{"message": "success"}, nil
			},
			requestMethod:  "GET",
			requestPath:    "/test",
			wantStatusCode: http.StatusOK,
			wantBody: HttpResponse{
				Data: map[string]interface{}{"message": "success"},
			},
		},
		{
			name: "should successfully register and handle endpoint with error response",
			setup: func(m mocks) {
				m.logMiddleware.EXPECT().LogMiddleware(mock.Anything).RunAndReturn(func(f constants.HttpFunction) constants.HttpFunction {
					return f
				})
			},
			method: "POST",
			path:   "/error",
			handlerFunc: func(ctx context.Context, w http.ResponseWriter, r *http.Request) (interface{}, error) {
				return nil, errors.New("something went wrong")
			},
			requestMethod:  "POST",
			requestPath:    "/error",
			wantStatusCode: http.StatusInternalServerError,
			wantBody: HttpResponse{
				Errors: []HttpErrorResponse{{Title: "something went wrong"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mocks{
				logMiddleware: httpmiddlewarelog_mocks.NewService(t),
				env:           env_mocks.NewService(t),
			}
			if tt.setup != nil {
				tt.setup(m)
			}
			s, _ := NewHttpServerService(context.Background(), m.logMiddleware, m.env)
			err := s.RegisterEndpoint(context.Background(), tt.method, tt.path, tt.handlerFunc)
			assert.NoError(t, err)

			req := httptest.NewRequest(tt.requestMethod, tt.requestPath, nil)
			w := httptest.NewRecorder()

			s.handler.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatusCode, w.Code)
			
			var gotBody HttpResponse
			err = json.Unmarshal(w.Body.Bytes(), &gotBody)
			assert.NoError(t, err)
			
			// We need to compare Data and Errors separately because of interface{} and map types
			if tt.wantBody.Data != nil {
				assert.Equal(t, tt.wantBody.Data, gotBody.Data)
			}
			if len(tt.wantBody.Errors) > 0 {
				assert.Equal(t, tt.wantBody.Errors, gotBody.Errors)
			}
		})
	}
}

func TestStart(t *testing.T) {
	type mocks struct {
		logMiddleware *httpmiddlewarelog_mocks.Service
		env           *env_mocks.Service
	}
	tests := []struct {
		name    string
		setup   func(m mocks)
		wantErr bool
	}{
		{
			name: "should successfully initialize server when port is provided",
			setup: func(m mocks) {
				m.env.EXPECT().Get(mock.Anything).Return(&env.Env{Port: "8080"})
			},
			wantErr: false,
		},
		{
			name: "should return error when port is empty",
			setup: func(m mocks) {
				m.env.EXPECT().Get(mock.Anything).Return(&env.Env{Port: ""})
			},
			wantErr: true,
		},
		{
			name: "should trigger error in goroutine when port is invalid",
			setup: func(m mocks) {
				m.env.EXPECT().Get(mock.Anything).Return(&env.Env{Port: "-1"})
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mocks{
				logMiddleware: httpmiddlewarelog_mocks.NewService(t),
				env:           env_mocks.NewService(t),
			}
			if tt.setup != nil {
				tt.setup(m)
			}
			s, _ := NewHttpServerService(context.Background(), m.logMiddleware, m.env)
			err := s.Start(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				assert.NotNil(t, s.server)
				if tt.name == "should trigger error in goroutine when port is invalid" {
					// wait for goroutine to fail
					time.Sleep(10 * time.Millisecond)
				} else {
					assert.Equal(t, ":8080", s.server.Addr)
				}
				// We don't wait for the goroutine to start Listening because it might fail on CI if port taken
				// But we should shutdown to be clean
				s.Shutdown(context.Background())
			}
		})
	}
}

func TestShutdown(t *testing.T) {
	type mocks struct {
		logMiddleware *httpmiddlewarelog_mocks.Service
		env           *env_mocks.Service
	}
	tests := []struct {
		name    string
		setup   func(m mocks, s *Implementation)
		wantErr bool
	}{
		{
			name: "should do nothing when server is nil",
			setup: func(m mocks, s *Implementation) {
				s.server = nil
			},
			wantErr: false,
		},
		{
			name: "should successfully shutdown server when initialized",
			setup: func(m mocks, s *Implementation) {
				m.env.EXPECT().Get(mock.Anything).Return(&env.Env{Port: "8081"})
				s.Start(context.Background())
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mocks{
				logMiddleware: httpmiddlewarelog_mocks.NewService(t),
				env:           env_mocks.NewService(t),
			}
			s, _ := NewHttpServerService(context.Background(), m.logMiddleware, m.env)
			if tt.setup != nil {
				tt.setup(m, s)
			}
			err := s.Shutdown(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("Shutdown() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
