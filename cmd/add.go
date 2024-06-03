package cmd

import (
	"log"

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
			log.Fatal(err)
			return
		}
		objectMap := internal.NewObjectMap()
		for _, entry := range entries {
			err := objectMap.AddObject(entry)
			if err != nil {
				log.Fatal(err)
				return
			}
		}
		for _, entry := range entries {
			obj, ok := objectMap.Get(entry.Path)
			if !ok {
				log.Fatalf("object not found: %s", entry.Path)
				return
			}
			err := internal.CreateFile(".bbgit/objects/"+obj.Hash, obj.Content)
			if err != nil {
				log.Fatal(err)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
