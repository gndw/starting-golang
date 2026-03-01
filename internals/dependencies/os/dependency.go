package os

import (
	"io"
	"io/fs"
	"os"
)

type Dependency interface {
	Stat(name string) (fs.FileInfo, error)
	Getenv(key string) string
	Stdout() io.Writer
	Exit(code int)
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

func (o *OSImpl) Stdout() io.Writer {
	return os.Stdout
}

func (o *OSImpl) Exit(code int) {
	os.Exit(code)
}
