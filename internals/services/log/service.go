package log

import (
	"context"
)

type Service interface {
	Info(ctx context.Context, format string, a ...any)
}
