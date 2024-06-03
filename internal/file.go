package internal

import (
	"os"
	"path/filepath"
)

type Entry struct {
	Path  string
	IsDir bool
}

func Explore(path string) ([]Entry, error) {
	var entries []Entry
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		entries = append(entries, Entry{Path: path, IsDir: info.IsDir()})
		return nil
	})
	if err != nil {
		return nil, err
	}
	leafToRoot := ReverseArray(entries)
	return leafToRoot, nil
}
