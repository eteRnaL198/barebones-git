package internal

import (
	"os"
	"path/filepath"
)

type File struct {
	Path  string
	IsDir bool
}

func Explore(path string) ([]File, error) {
	var paths []File
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		paths = append(paths, File{Path: path, IsDir: info.IsDir()})
		return nil
	})
	if err != nil {
		return nil, err
	}
	leafToRoot := ReverseArray(paths)
	return leafToRoot, nil
}
