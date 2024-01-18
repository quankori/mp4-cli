package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "Split file mp4",
	Long:  `Split file mp4 in the directory.`,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing files:")
		split(args)
	},
}

func splitVideo(inputFile, startTime, endTime, outputFile string) error {
	cmd := exec.Command("ffmpeg", "-i", inputFile, "-ss", startTime, "-to", endTime, "-c", "copy", outputFile)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg error: %w", err)
	}

	return nil
}

func split(args []string) {
	outputFile := "output.mp4"
	if len(os.Args) < 4 {
		fmt.Println("Usage: mp4-cli <filename> <timestart> <timeend>")
		os.Exit(1)
	}

	filename := args[0]
	timeStart := args[1]
	timeEnd := args[2]
	if err := splitVideo(filename, timeStart, timeEnd, outputFile); err != nil {
		fmt.Println("Error splitting video:", err)
	} else {
		fmt.Println("Video split successfully")
	}
}
