package env

import (
	"context"
)

type Env struct {
	Port string
}

type Service interface {
	Get(ctx context.Context) *Env
}
