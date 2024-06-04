package internal

import (
	"os"
	"path/filepath"
)

type File struct {
	Path    string
	Content string
}

func NewFile(path, content string) *File {
	return &File{
		Path:    path,
		Content: content,
	}
}

func CreateFile(file File) error {
	content := []byte(file.Content)
	err := os.WriteFile(file.Path, content, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ExploreFiles(path string) ([]File, error) {
	var files []File
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		contentInBytes, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		files = append(files, File{Path: path, Content: string(contentInBytes)})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
