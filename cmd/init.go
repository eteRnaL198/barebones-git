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
	gitDir := filepath.Join(".bbgit")
	err := os.Mkdir(gitDir, 0755)
	if err != nil {
		return err
	}

	objectsDir := filepath.Join(gitDir, "objects")
	err = os.Mkdir(objectsDir, 0755)
	if err != nil {
		return err
	}

	return nil
}
