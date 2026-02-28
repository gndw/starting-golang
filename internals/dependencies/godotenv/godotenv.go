package godotenv

import "github.com/joho/godotenv"

//mockery:generate: true
type Dependency interface {
	Load(filenames ...string) error
	Overload(filenames ...string) error
}

type GodotenvImpl struct{}

func NewGodotenv() *GodotenvImpl {
	return &GodotenvImpl{}
}

func (g *GodotenvImpl) Load(filenames ...string) error {
	return godotenv.Load(filenames...)
}

func (g *GodotenvImpl) Overload(filenames ...string) error {
	return godotenv.Overload(filenames...)
}
