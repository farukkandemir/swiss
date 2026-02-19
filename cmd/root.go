package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "swiss",
	Short: "A personal toolkit for getting things done",
	Long:  "Swiss is a collection of handy utilities and commands, built to streamline your workflow from the terminal.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Swiss Use --help to see available commands")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
