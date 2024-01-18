package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mp4-cli",
	Short: "My CLI Application",
	Long:  `This is my CLI application using Viper and Cobra.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(listCmd)
}
