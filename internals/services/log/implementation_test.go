package log

import (
	"context"
	"testing"

	os_mocks "github.com/gndw/starting-golang/internals/dependencies/os/mocks"
	slog_mocks "github.com/gndw/starting-golang/internals/dependencies/slog/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewLogService(t *testing.T) {
	type mocks struct {
		slogDependency *slog_mocks.Dependency
		osDependency   *os_mocks.Dependency
		logger         *slog_mocks.Logger
	}
	tests := []struct {
		name    string
		setup   func(m mocks)
		wantErr error
	}{
		{
			name: "should successfully create log service when dependencies are provided",
			setup: func(m mocks) {
				m.osDependency.EXPECT().Stdout().Return(nil)
				m.slogDependency.EXPECT().NewJSONLogger(mock.Anything).Return(m.logger)
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mocks{
				slogDependency: slog_mocks.NewDependency(t),
				osDependency:   os_mocks.NewDependency(t),
				logger:         slog_mocks.NewLogger(t),
			}
			if tt.setup != nil {
				tt.setup(m)
			}

			got, err := NewLogService(context.Background(), m.slogDependency, m.osDependency)

			assert.NoError(t, err)
			assert.NotNil(t, got)
			impl := got.(*Implementation)
			assert.Equal(t, m.logger, impl.logger)
			assert.Equal(t, m.osDependency, impl.os)
		})
	}
}

func TestImplementation_LogMethods(t *testing.T) {
	type mocks struct {
		logger *slog_mocks.Logger
		os     *os_mocks.Dependency
	}
	ctx := context.Background()
	msg := "test message"
	args := []any{"key", "value"}

	tests := []struct {
		name   string
		action func(s *Implementation, ctx context.Context, msg string, args ...any)
		setup  func(m mocks)
	}{
		{
			name: "should call DebugContext when Debug is called",
			action: func(s *Implementation, ctx context.Context, msg string, args ...any) {
				s.Debug(ctx, msg, args...)
			},
			setup: func(m mocks) {
				m.logger.EXPECT().DebugContext(ctx, msg, args).Return()
			},
		},
		{
			name: "should call InfoContext when Info is called",
			action: func(s *Implementation, ctx context.Context, msg string, args ...any) {
				s.Info(ctx, msg, args...)
			},
			setup: func(m mocks) {
				m.logger.EXPECT().InfoContext(ctx, msg, args).Return()
			},
		},
		{
			name: "should call WarnContext when Warn is called",
			action: func(s *Implementation, ctx context.Context, msg string, args ...any) {
				s.Warn(ctx, msg, args...)
			},
			setup: func(m mocks) {
				m.logger.EXPECT().WarnContext(ctx, msg, args).Return()
			},
		},
		{
			name: "should call ErrorContext when Error is called",
			action: func(s *Implementation, ctx context.Context, msg string, args ...any) {
				s.Error(ctx, msg, args...)
			},
			setup: func(m mocks) {
				m.logger.EXPECT().ErrorContext(ctx, msg, args).Return()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mocks{
				logger: slog_mocks.NewLogger(t),
				os:     os_mocks.NewDependency(t),
			}
			if tt.setup != nil {
				tt.setup(m)
			}
			s := &Implementation{
				logger: m.logger,
				os:     m.os,
			}
			tt.action(s, ctx, msg, args...)
		})
	}
}

func TestImplementation_Fatal(t *testing.T) {
	type mocks struct {
		logger *slog_mocks.Logger
		os     *os_mocks.Dependency
	}
	ctx := context.Background()
	msg := "fatal message"
	args := []any{"key", "value"}

	tests := []struct {
		name  string
		setup func(m mocks)
	}{
		{
			name: "should log error and exit when Fatal is called",
			setup: func(m mocks) {
				m.logger.EXPECT().ErrorContext(ctx, msg, args).Return()
				m.os.EXPECT().Exit(1).Return()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mocks{
				logger: slog_mocks.NewLogger(t),
				os:     os_mocks.NewDependency(t),
			}
			if tt.setup != nil {
				tt.setup(m)
			}
			s := &Implementation{
				logger: m.logger,
				os:     m.os,
			}
			s.Fatal(ctx, msg, args...)
		})
	}
}
