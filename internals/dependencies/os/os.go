package os

import (
	"io/fs"
	"os"
)

//go:generate mockery --name Dependency
type Dependency interface {
	Stat(name string) (fs.FileInfo, error)
	Getenv(key string) string
}

type OSImpl struct{}

func NewOS() *OSImpl {
	return &OSImpl{}
}

func (o *OSImpl) Stat(name string) (fs.FileInfo, error) {
	return os.Stat(name)
}

func (o *OSImpl) Getenv(key string) string {
	return os.Getenv(key)
}
