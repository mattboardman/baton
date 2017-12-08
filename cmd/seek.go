package cmd

import (
	"fmt"
	"strconv"

	"github.com/joshuathompson/baton/api"
	"github.com/spf13/cobra"
)

func seekToPosition(cmd *cobra.Command, args []string) {
	pos, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Printf("Time specified could not be converted to seconds")
		return
	}

	fmt.Println(pos)

	err = api.SeekToPosition(pos * 1000)

	if err != nil {
		fmt.Printf("Failed to restart current track\n")
	} else {
		fmt.Printf("Replaying current track\n")
	}
}

func init() {
	rootCmd.AddCommand(seekCmd)
}

var seekCmd = &cobra.Command{
	Use:   "seek [pos]",
	Short: "Skip to a specific time (seconds) of the current song",
	Long:  `Skip to a specific time (seconds) of the current song`,
	Run:   seekToPosition,
}