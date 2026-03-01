package slog

import (
	"context"
	"io"
	"log/slog"
)

type Dependency interface {
	NewJSONLogger(w io.Writer) Logger
}

type Logger interface {
	DebugContext(ctx context.Context, msg string, args ...any)
	InfoContext(ctx context.Context, msg string, args ...any)
	WarnContext(ctx context.Context, msg string, args ...any)
	ErrorContext(ctx context.Context, msg string, args ...any)
}

type SlogImpl struct{}

func NewSlog() *SlogImpl {
	return &SlogImpl{}
}

func (s *SlogImpl) NewJSONLogger(w io.Writer) Logger {
	return slog.New(slog.NewJSONHandler(w, nil))
}
