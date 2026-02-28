package env

import (
	"context"

	"github.com/gndw/starting-golang/internals/dependencies/godotenv"
	"github.com/gndw/starting-golang/internals/dependencies/os"
)

type Implementation struct {
	env *Env
}

func NewEnvService(ctx context.Context, godotenv godotenv.Dependency, os os.Dependency) (*Implementation, error) {
	// load .env file from root
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	// load .local.env if it exists
	if _, err := os.Stat(".local.env"); err == nil {
		_ = godotenv.Overload(".local.env")
	}

	return &Implementation{
		env: &Env{
			Port: os.Getenv("PORT"),
		},
	}, nil
}

func (m *Implementation) Get(ctx context.Context) *Env {
	return m.env
}
