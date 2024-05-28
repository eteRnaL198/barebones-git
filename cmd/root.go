package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bbgit",
	Short: "BarebonesGit is a simple Git clone",
	Long:  `BareBonesGit is a small-scale clone of Git, written in Go.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you can define additional commands and flags.
}
