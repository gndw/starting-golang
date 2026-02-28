package env

import (
	"context"
)

type Env struct {
	Port string
}

//go:generate mockery --name Service
type Service interface {
	Get(ctx context.Context) *Env
}
