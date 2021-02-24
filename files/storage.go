package files

import "io"

// Storage defines the behavior for file operations
// Implementation exists for localdisk but cloud storage is tbd
type Storage interface {
	Save(path string, file io.Reader) error
}
