package filesystem

import (
	"os"
	"path/filepath"
)

// fileFinder satisfies the Finder interface
type fileFinder struct{}

// NewFileFinder is the factory function for the file finder
func NewFileFinder() Finder {
	return &fileFinder{}
}

// Find ...
func (fm *fileFinder) Find(path string) ([]string, error) {
	return filepath.Glob(path)
}

// Read ...
func (fm *fileFinder) Read(match string) ([]byte, error) {
	return os.ReadFile(match)
}
