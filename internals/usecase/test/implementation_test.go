package test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/gndw/starting-golang/internals/models"
	"github.com/gndw/starting-golang/internals/repositories/inmemorydb"
	inMemoryDbRepositoryMocks "github.com/gndw/starting-golang/internals/repositories/inmemorydb/mocks"
	"github.com/stretchr/testify/mock"
)

func Test(t *testing.T) {
	type args struct {
		ctx     context.Context
		request models.TestRequest
	}
	type fields struct {
		inMemoryDbRepository *inMemoryDbRepositoryMocks.Repository
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		mocks        func(fields *fields)
		wantResponse models.TestResponse
		wantErr      bool
	}{
		{
			name:   "success",
			fields: fields{inMemoryDbRepository: inMemoryDbRepositoryMocks.NewRepository(t)},
			args: args{
				ctx:     context.Background(),
				request: models.TestRequest{UserID: 100},
			},
			mocks: func(fields *fields) {
				fields.inMemoryDbRepository.On("GetUserData", context.Background(), int64(100)).
					Return(models.User{ID: 100, FullName: "John Doe"}, nil)
			},
			wantResponse: models.TestResponse{UserID: 100, FullName: "John Doe"},
		},
		{
			name:   "failed when inMemoryDb returning error",
			fields: fields{inMemoryDbRepository: inMemoryDbRepositoryMocks.NewRepository(t)},
			args: args{
				ctx:     context.Background(),
				request: models.TestRequest{UserID: 100},
			},
			mocks: func(fields *fields) {
				fields.inMemoryDbRepository.On("GetUserData", mock.Anything, mock.Anything).
					Return(models.User{}, errors.New("failed"))
			},
			wantResponse: models.TestResponse{},
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocks(&tt.fields)
			m := &Implementation{
				inMemoryDbRepository: tt.fields.inMemoryDbRepository,
			}
			gotResponse, err := m.Test(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Implementation.Test() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("Implementation.Test() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestNewUsecase(t *testing.T) {
	type args struct {
		ctx                  context.Context
		inMemoryDbRepository inmemorydb.Repository
	}
	targs := args{
		ctx:                  context.Background(),
		inMemoryDbRepository: inMemoryDbRepositoryMocks.NewRepository(t),
	}
	tests := []struct {
		name    string
		args    args
		want    *Implementation
		wantErr bool
	}{
		{
			name: "success",
			args: targs,
			want: &Implementation{
				inMemoryDbRepository: targs.inMemoryDbRepository,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUsecase(tt.args.ctx, tt.args.inMemoryDbRepository)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUsecase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
