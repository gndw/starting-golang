package test

import (
	"context"
	"net/http"
)

type Handler interface {
	Test(ctx context.Context, rw http.ResponseWriter, r *http.Request) (data interface{}, err error)
}
