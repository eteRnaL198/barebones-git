package cmd

import (
	"log"
	"os"

	"github.com/eteRnaL198/barebones-git/internal"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Record changes to the repository",
	Run: func(cmd *cobra.Command, args []string) {
		err := commit()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}

func commit() error {
	indexContentsInBytes, err := os.ReadFile(".bbgit/index")
	if err != nil {
		return nil
	}
	indexContents := internal.ParseIndexFile(string(indexContentsInBytes))
	tree := internal.ParseIndexContents(indexContents)
	rootTreeHash := internal.CreateTreeObject(*tree)

	commitContent := "tree " + rootTreeHash
	headInBytes, err := os.ReadFile(".bbgit/HEAD")
	if err == nil {
		commitContent += "\nparent " + string(headInBytes)
	}
	commitHash := internal.CalcHash(commitContent)
	internal.CreateFile(*internal.NewFile(".bbgit/objects/"+commitHash, commitContent))

	headFile, err := os.OpenFile(".bbgit/HEAD", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer headFile.Close()
	_, err = headFile.Write([]byte(commitHash))
	if err != nil {
		return err
	}

	return nil
}
