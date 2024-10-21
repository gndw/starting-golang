package resources

import (
	"context"

	testHandler "github.com/gndw/starting-golang/internals/handlers/test"
	"github.com/gndw/starting-golang/internals/repositories/inmemorydb"
	"github.com/gndw/starting-golang/internals/services/httpserver"
	testUsecase "github.com/gndw/starting-golang/internals/usecase/test"
)

type Resource struct {
	HttpService httpserver.Service
}

// Init will initialize all services and layers

func Init(ctx context.Context) (resource Resource, err error) {

	httpService, err := httpserver.NewHttpService(ctx)
	if err != nil {
		return resource, err
	}

	inMemoryDbRepository, err := inmemorydb.NewRepository(ctx)
	if err != nil {
		return resource, err
	}

	testUsecase, err := testUsecase.NewUsecase(ctx, inMemoryDbRepository)
	if err != nil {
		return resource, err
	}

	_, err = testHandler.NewHandler(ctx, httpService, testUsecase)
	if err != nil {
		return resource, err
	}

	return Resource{
		HttpService: httpService,
	}, nil
}
