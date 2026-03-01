package resources

import (
	"context"

	"github.com/gndw/starting-golang/internals/dependencies/godotenv"
	"github.com/gndw/starting-golang/internals/dependencies/os"
	"github.com/gndw/starting-golang/internals/dependencies/slog"
	testHandler "github.com/gndw/starting-golang/internals/handlers/test"
	"github.com/gndw/starting-golang/internals/repositories/inmemorydb"
	"github.com/gndw/starting-golang/internals/services/env"
	"github.com/gndw/starting-golang/internals/services/httpmiddlewarelog"
	"github.com/gndw/starting-golang/internals/services/httpserver"
	"github.com/gndw/starting-golang/internals/services/log"
	testUsecase "github.com/gndw/starting-golang/internals/usecase/test"
)

type Resource struct {
	HttpServerService httpserver.Service
}

// Init will initialize all services and layers

func Init(ctx context.Context) (resource Resource, err error) {

	osDependency := os.NewOS()
	godotenvDependency := godotenv.NewGodotenv()
	slogDependency := slog.NewSlog()

	logService, err := log.NewLogService(ctx, slogDependency, osDependency)
	if err != nil {
		return resource, err
	}

	envService, err := env.NewEnvService(ctx, godotenvDependency, osDependency)
	if err != nil {
		return resource, err
	}

	httpMiddlewareService, err := httpmiddlewarelog.NewHttpMiddlewareService(ctx, logService)
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
