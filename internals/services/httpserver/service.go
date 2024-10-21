package httpserver

import (
	"context"

	"github.com/gndw/starting-golang/internals/constants"
)

//go:generate mockery --name Service
type Service interface {
	RegisterEndpoint(ctx context.Context, method string, path string, f constants.HttpFunction) error
	Serve(ctx context.Context) error
}
