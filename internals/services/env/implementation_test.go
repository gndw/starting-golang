package env_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	godotenv_mocks "github.com/gndw/starting-golang/internals/dependencies/godotenv/mocks"
	os_mocks "github.com/gndw/starting-golang/internals/dependencies/os/mocks"
	"github.com/gndw/starting-golang/internals/services/env"
)

func TestNewEnvService(t *testing.T) {
	type mocks struct {
		godotenv *godotenv_mocks.Dependency
		os       *os_mocks.Dependency
	}
	errLoad := errors.New("load error")
	tests := []struct {
		name    string
		setup   func(m mocks)
		want    *env.Env
		wantErr error
	}{
		{
			name: "success only .env",
			setup: func(m mocks) {
				m.godotenv.EXPECT().Load([]string{".env"}).Return(nil)
				m.os.EXPECT().Stat(".local.env").Return(nil, errors.New("not found"))
				m.os.EXPECT().Getenv("PORT").Return("8080")
			},
			want:    &env.Env{Port: "8080"},
			wantErr: nil,
		},
		{
			name: "success with .local.env",
			setup: func(m mocks) {
				m.godotenv.EXPECT().Load([]string{".env"}).Return(nil)
				m.os.EXPECT().Stat(".local.env").Return(nil, nil)
				m.godotenv.EXPECT().Overload([]string{".local.env"}).Return(nil)
				m.os.EXPECT().Getenv("PORT").Return("9090")
			},
			want:    &env.Env{Port: "9090"},
			wantErr: nil,
		},
		{
			name: "fail load .env",
			setup: func(m mocks) {
				m.godotenv.EXPECT().Load([]string{".env"}).Return(errLoad)
			},
			wantErr: errLoad,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mocks{
				godotenv: godotenv_mocks.NewDependency(t),
				os:       os_mocks.NewDependency(t),
			}
			if tt.setup != nil {
				tt.setup(m)
			}

			got, gotErr := env.NewEnvService(context.Background(), m.godotenv, m.os)
			if !errors.Is(gotErr, tt.wantErr) {
				t.Errorf("NewEnvService() error = %v, wantErr %v", gotErr, tt.wantErr)
				return
			}
			if gotErr != nil {
				return
			}

			if !reflect.DeepEqual(got.Get(context.Background()), tt.want) {
				t.Errorf("NewEnvService() = %v, want %v", got.Get(context.Background()), tt.want)
			}
		})
	}
}
