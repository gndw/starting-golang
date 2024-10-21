package test

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/gndw/starting-golang/internals/models"
	httpserverServiceMocks "github.com/gndw/starting-golang/internals/services/httpserver/mocks"
	testUsecaseMocks "github.com/gndw/starting-golang/internals/usecase/test/mocks"

	"github.com/stretchr/testify/mock"
)

func TestImplementation_Test(t *testing.T) {
	type args struct {
		ctx context.Context
		rw  http.ResponseWriter
		r   *http.Request
	}
	type fields struct {
		testUsecase *testUsecaseMocks.Usecase
	}
	tests := []struct {
		name     string
		args     args
		fields   fields
		mocks    func(fields *fields)
		wantData interface{}
		wantErr  bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				r: func() *http.Request {
					r, _ := http.NewRequest("", "/test", strings.NewReader(`{"user_id":100}`))
					return r
				}(),
			},
			fields: fields{testUsecase: testUsecaseMocks.NewUsecase(t)},
			mocks: func(fields *fields) {
				fields.testUsecase.On("Test", context.Background(), models.TestRequest{UserID: 100}).
					Return(models.TestResponse{UserID: 100, FullName: "John Doe"}, nil)
			},
			wantData: models.TestResponse{UserID: 100, FullName: "John Doe"},
		},
		{
			name: "failed when test usecase is returning failed",
			args: args{
				ctx: context.Background(),
				r: func() *http.Request {
					r, _ := http.NewRequest("", "/test", strings.NewReader(`{"user_id":100}`))
					return r
				}(),
			},
			fields: fields{testUsecase: testUsecaseMocks.NewUsecase(t)},
			mocks: func(fields *fields) {
				fields.testUsecase.On("Test", mock.Anything, mock.Anything).Return(models.TestResponse{}, errors.New("failed!"))
			},
			wantData: nil,
			wantErr:  true,
		},
		{
			name: "failed when body parsing is returning error",
			args: args{
				ctx: context.Background(),
				r: func() *http.Request {
					r, _ := http.NewRequest("", "/test", strings.NewReader(`xxx`))
					return r
				}(),
			},
			fields:   fields{testUsecase: testUsecaseMocks.NewUsecase(t)},
			mocks:    func(fields *fields) {},
			wantData: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocks(&tt.fields)
			m := &Implementation{
				testUsecase: tt.fields.testUsecase,
			}
			gotData, err := m.Test(tt.args.ctx, tt.args.rw, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Implementation.Test() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("Implementation.Test() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func TestNewHandler(t *testing.T) {
	type args struct {
		ctx         context.Context
		httpService *httpserverServiceMocks.Service
		testUsecase *testUsecaseMocks.Usecase
	}
	targs := args{
		ctx:         context.Background(),
		httpService: httpserverServiceMocks.NewService(t),
		testUsecase: testUsecaseMocks.NewUsecase(t),
	}
	tests := []struct {
		name    string
		args    args
		mocks   func(args *args)
		want    *Implementation
		wantErr bool
	}{
		{
			name: "success",
			args: targs,
			mocks: func(args *args) {
				args.httpService.On("RegisterEndpoint", context.Background(), http.MethodPost, "/test", mock.Anything).Return(nil)
			},
			want: &Implementation{
				testUsecase: targs.testUsecase,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocks(&tt.args)
			got, err := NewHandler(tt.args.ctx, tt.args.httpService, tt.args.testUsecase)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
