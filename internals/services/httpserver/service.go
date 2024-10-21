package httpserver

import (
	"context"
	"net/http"
)

type HttpFunction func(ctx context.Context, w http.ResponseWriter, r *http.Request) (interface{}, error)

//go:generate mockery --name Service
type Service interface {
	RegisterEndpoint(ctx context.Context, method string, path string, f HttpFunction) error
	Serve(ctx context.Context) error
}
