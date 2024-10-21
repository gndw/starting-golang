package constants

import (
	"context"
	"net/http"
)

type HttpFunction func(ctx context.Context, w http.ResponseWriter, r *http.Request) (interface{}, error)
