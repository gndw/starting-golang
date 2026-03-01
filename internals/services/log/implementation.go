package log

import (
	"context"

	"github.com/gndw/starting-golang/internals/dependencies/os"
	"github.com/gndw/starting-golang/internals/dependencies/slog"
)

type Implementation struct {
	logger slog.Logger
	os     os.Dependency
}

func NewLogService(ctx context.Context, slogDependency slog.Dependency, osDependency os.Dependency) (*Implementation, error) {
	logger := slogDependency.NewJSONLogger(osDependency.Stdout())
	return &Implementation{
		logger: logger,
		os:     osDependency,
	}, nil
}

func (m *Implementation) Debug(ctx context.Context, msg string, args ...any) {
	m.logger.DebugContext(ctx, msg, args...)
}

func (m *Implementation) Info(ctx context.Context, msg string, args ...any) {
	m.logger.InfoContext(ctx, msg, args...)
}

func (m *Implementation) Warn(ctx context.Context, msg string, args ...any) {
	m.logger.WarnContext(ctx, msg, args...)
}

func (m *Implementation) Error(ctx context.Context, msg string, args ...any) {
	m.logger.ErrorContext(ctx, msg, args...)
}

func (m *Implementation) Fatal(ctx context.Context, msg string, args ...any) {
	m.logger.ErrorContext(ctx, msg, args...)
	m.os.Exit(1)
}
