package httpserver

import (
	"context"

	"github.com/gndw/starting-golang/internals/constants"
)

type Service interface {
	RegisterEndpoint(ctx context.Context, method string, path string, f constants.HttpFunction) error
	Start(ctx context.Context) error
	Shutdown(ctx context.Context) error
}
