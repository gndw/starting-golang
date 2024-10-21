package log

import (
	"context"
)

//go:generate mockery --name Service
type Service interface {
	Info(ctx context.Context, format string, a ...any)
}
