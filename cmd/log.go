package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show commit logs",
	Run: func(cmd *cobra.Command, args []string) {
		err := displayCommitLog()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}

func displayCommitLog() error {
	logFile, err := os.ReadFile(".bbgit/logs")
	if err != nil {
		return err
	}
	fmt.Println(string(logFile))
	return nil
}
