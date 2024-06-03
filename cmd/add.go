package cmd

import (
	"fmt"

	"github.com/eteRnaL198/barebones-git/internal"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <filename>",
	Short: "Add file contents to the index",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		files, err := internal.Explore(filePath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		for _, file := range files {
			if file.IsDir {
				fmt.Println("create tree for:", file.Path)
			} else {
				fmt.Println("create blob for:", file.Path)
			}
		}

		// err := addObjectsToIndex()
		// if err != nil {
		// 	fmt.Println("Error:", err)
		// } else {
		// 	fmt.Println("Changes added to index.")
		// }
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addObjectsToIndex(filePath string) error {
	hash, err := internal.CalcHash(filePath)
	if err != nil {
		return err
	}
	err = internal.CreateBlob(filePath, ".bbgit/objects/"+hash)
	if err != nil {
		return err
	}
	return nil
}
