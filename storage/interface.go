package storage

import (
	"io"
	"time"
)

type Storage interface {
	Stat(name string) (FileInfo, error)
	ReadDir(dirname string) ([]FileInfo, error)
	Create(name string) (io.WriteCloser, error)
	Open(name string) (io.ReadCloser, error)
	Remove(name string) error
	RemoveAll(path string) error
	Rename(oldpath, newpath string) error
	Mkdir(name string) error
	MkdirAll(path string) error
}

type FileInfo interface {
	Name() string       // base name of the file
	Size() int64        // length in bytes for regular files; system-dependent for others
	ModTime() time.Time // modification time
	IsDir() bool
}
