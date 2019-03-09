package storage

import (
	"io"
	"io/ioutil"
	"os"
)

type local struct{}

// NewLocal ...
func NewLocal() Storage {
	return local{}
}

func (local) Stat(name string) (FileInfo, error) {
	return os.Stat(name)
}

func (local) ReadDir(dirname string) ([]FileInfo, error) {
	fi, err := ioutil.ReadDir(dirname)
	f := make([]FileInfo, len(fi))
	for i := range fi {
		f[i] = fi[i]
	}
	return f, err
}

func (local) Create(name string) (io.WriteCloser, error) {
	return os.Create(name)
}

func (local) Open(name string) (io.ReadCloser, error) {
	return os.Open(name)
}

func (local) Remove(name string) error {
	return os.Remove(name)
}

func (local) RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func (local) Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}

func (local) Mkdir(name string) error {
	return os.Mkdir(name, 0755)
}

func (local) MkdirAll(path string) error {
	return os.MkdirAll(path, 0755)
}
