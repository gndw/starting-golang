package log

import (
	"context"
	"fmt"
)

type Implementation struct {
}

func NewLogService(ctx context.Context) (*Implementation, error) {
	return &Implementation{}, nil
}

func (m *Implementation) Info(ctx context.Context, format string, a ...any) {
	fmt.Println(fmt.Sprintf(format, a...))
}
