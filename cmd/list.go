package cmd

import (
	"fmt"
	"os/exec"

	"github.com/quankori/mp4-cli/internal/utils"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List files mp4",
	Long:  `List all files mp4 in the directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing files:")
		utils.ListFiles()
	},
}

func init() {
	// Here you can define flags and configuration settings.
}

func isFFmpegInstalled() bool {
	cmd := exec.Command("ffmpeg", "-version")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
