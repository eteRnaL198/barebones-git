package cmd

import (
	"log"
	"path/filepath"

	"github.com/eteRnaL198/barebones-git/internal"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add file contents to the index",
	Run: func(cmd *cobra.Command, args []string) {
		err := add(filepath.Join("root"))
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func add(filePath string) error {
	files, err := internal.ExploreFiles(filePath)
	if err != nil {
		return err
	}

	var blobs []blob
	for _, file := range files {
		blob := NewBlob(file.Content, file.Path)
		blobs = append(blobs, *blob)
	}

	var indexContent string
	for _, blob := range blobs {
		err := internal.CreateFile(*internal.NewFile(".bbgit/objects/"+blob.Hash, blob.Content))
		if err != nil {
			return err
		}
		indexContent += blob.Hash + " " + blob.Path + "\n"
	}
	err = internal.CreateFile(*internal.NewFile(".bbgit/index", indexContent))
	if err != nil {
		return err
	}

	return nil
}

type blob struct {
	Path    string
	Hash    string
	Content string
}

func NewBlob(content, path string) *blob {
	return &blob{
		Path:    path,
		Hash:    internal.CalcHash(content),
		Content: "blob\n" + content,
	}
}
