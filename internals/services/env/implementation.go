package env

import (
	"context"
	"os"

	"github.com/joho/godotenv"
)

type Implementation struct {
	env *Env
}

func NewEnvService(ctx context.Context) (*Implementation, error) {
	// load .env file from root
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
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
