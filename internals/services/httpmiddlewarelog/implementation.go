package httpmiddlewarelog

import (
	"context"

	"github.com/gndw/starting-golang/internals/services/log"
)

type Implementation struct {
	logService log.Service
}

func NewHttpMiddlewareService(ctx context.Context, logService log.Service) (*Implementation, error) {
	return &Implementation{logService: logService}, nil
}
