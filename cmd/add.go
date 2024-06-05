package cmd

import (
	"log"
	"os"

	"github.com/eteRnaL198/barebones-git/internal"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <filename>",
	Short: "Add file contents to the index",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := add(args[0])
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

	for _, blob := range blobs {
		err := internal.CreateFile(*internal.NewFile(".bbgit/objects/"+blob.Hash, blob.Content))
		if err != nil {
			return err
		}

		indexFile, err := os.OpenFile(".bbgit/index", os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			return err
		}
		defer indexFile.Close()
		_, err = indexFile.WriteString(blob.Hash + " " + blob.Path + "\n")
		if err != nil {
			return err
		}
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
