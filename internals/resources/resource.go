package resources

import (
	"context"

	testHandler "github.com/gndw/starting-golang/internals/handlers/test"
	"github.com/gndw/starting-golang/internals/repositories/inmemorydb"
	"github.com/gndw/starting-golang/internals/services/env"
	"github.com/gndw/starting-golang/internals/services/httpmiddleware"
	"github.com/gndw/starting-golang/internals/services/httpserver"
	"github.com/gndw/starting-golang/internals/services/log"
	testUsecase "github.com/gndw/starting-golang/internals/usecase/test"
)

type Resource struct {
	HttpServerService httpserver.Service
}

// Init will initialize all services and layers

func Init(ctx context.Context) (resource Resource, err error) {

	logService, err := log.NewLogService(ctx)
	if err != nil {
		return resource, err
	}

	envService, err := env.NewEnvService(ctx)
	if err != nil {
		return resource, err
	}

	httpMiddlewareService, err := httpmiddleware.NewHttpMiddlewareService(ctx, logService)
	if err != nil {
		return resource, err
	}

	httpServerService, err := httpserver.NewHttpServerService(ctx, httpMiddlewareService, envService)
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

	_, err = testHandler.NewHandler(ctx, httpServerService, testUsecase)
	if err != nil {
		return resource, err
	}

	return Resource{
		HttpServerService: httpServerService,
	}, nil
}
