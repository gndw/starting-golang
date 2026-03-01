package test

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gndw/starting-golang/internals/models"
	"github.com/gndw/starting-golang/internals/services/httpserver"
	"github.com/gndw/starting-golang/internals/usecase/test"
)

type Implementation struct {
	testUsecase test.Usecase
}

func NewHandler(ctx context.Context, httpService httpserver.Service, testUsecase test.Usecase) (Handler, error) {
	h := &Implementation{
		testUsecase: testUsecase,
	}
	httpService.RegisterEndpoint(ctx, http.MethodPost, "/test", h.Test)
	return h, nil
}

func (m *Implementation) Test(ctx context.Context, rw http.ResponseWriter, r *http.Request) (data interface{}, err error) {

	var testRequest models.TestRequest

	err = json.NewDecoder(r.Body).Decode(&testRequest)
	if err != nil {
		return nil, err
	}

	disburseResponse, err := m.testUsecase.Test(ctx, testRequest)
	if err != nil {
		return nil, err
	}

	return disburseResponse, nil
}
