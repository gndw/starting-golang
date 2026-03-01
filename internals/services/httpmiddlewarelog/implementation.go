package httpmiddlewarelog

import (
	"context"
	"net/http"

	"github.com/gndw/starting-golang/internals/constants"
	"github.com/gndw/starting-golang/internals/services/log"
)

type Implementation struct {
	logService log.Service
}

func NewHttpMiddlewareService(ctx context.Context, logService log.Service) (*Implementation, error) {
	return &Implementation{logService: logService}, nil
}

func (m *Implementation) LogMiddleware(f constants.HttpFunction) constants.HttpFunction {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) (interface{}, error) {

		resp, err := f(ctx, w, r)

		// do something else with the log
		m.logService.Info(ctx, "[incoming-http]...")

		return resp, err
	}
}
