package cmd

import (
	"fmt"
	"os"

	"github.com/eteRnaL198/barebones-git/internal"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <filename>",
	Short: "Add file contents to the index",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		entries, err := internal.Explore(filePath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		for _, entry := range entries {
			if entry.IsDir {
				fmt.Println("create tree for:", entry.Path)
				innerFiles, err := os.ReadDir(entry.Path)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				for _, innerFile := range innerFiles {
					fmt.Println(entry.Path+"has file:", innerFile.Name())
				}
			} else {
				fmt.Println("create blob for:", entry.Path)
				// err := addObjectsToIndex()
				// if err != nil {
				// 	fmt.Println("Error:", err)
				// } else {
				// 	fmt.Println("Changes added to index.")
				// }
			}
		}

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
