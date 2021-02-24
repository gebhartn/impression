package files

import (
	"io"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

// Local is an implementation of the store interface for localdisk
type Local struct {
	max      int
	basePath string
}

// NewLocal creates a new Local filesystem at the path specified
func NewLocal(basePath string, max int) (*Local, error) {
	p, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}

	return &Local{basePath: p, max: max}, nil
}

// Save the contexts of the Writer
func (l *Local) Save(path string, contents io.Reader) error {
	p := l.fullPath(path)
	d := filepath.Dir(p)

	err := os.MkdirAll(d, os.ModePerm)
	if err != nil {
		return xerrors.Errorf("Unable to create directory: %w", err)
	}

	_, err = os.Stat(p)
	if err == nil {
		err = os.Remove(p)
		if err != nil {
			return xerrors.Errorf("Unable to delete file: %w", err)
		}
	} else if !os.IsNotExist(err) {
		return xerrors.Errorf("Unable to get file info: %w", err)
	}

	f, err := os.Create(p)
	if err != nil {
		return xerrors.Errorf("Unable to create file: %w", err)
	}
	defer f.Close()

	_, err = io.Copy(f, contents)
	if err != nil {
		return xerrors.Errorf("Unable to write to file: %w", err)
	}

	return nil
}

func (l *Local) fullPath(path string) string {
	return filepath.Join(l.basePath, path)
}

// Get the file at a specified path
func (l *Local) Get(path string) string {
	return filepath.Join(l.basePath, path)
}
