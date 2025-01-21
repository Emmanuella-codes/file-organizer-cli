package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "file-organizer",
	Short: "",
	Long: "File Organizer is a simple tool to organize your files and folders.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the File Organizer CLI!")
	},
}

func Execute() error {
	return rootCmd.Execute()
}