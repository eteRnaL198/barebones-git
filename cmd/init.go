package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new, empty repository",
	Run: func(cmd *cobra.Command, args []string) {
		err := initRepo()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Initialized empty BarebonesGit repository.")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initRepo() error {
	gitDirPath := filepath.Join(".bbgit")
	err := os.Mkdir(gitDirPath, 0755)
	if err != nil {
		return err
	}

	objectsDirPath := filepath.Join(gitDirPath, "objects")
	err = os.Mkdir(objectsDirPath, 0755)
	if err != nil {
		return err
	}

	indexFilePath := filepath.Join(gitDirPath, "index")
	_, err = os.Create(indexFilePath)
	if err != nil {
		return err
	}

	return nil
}
