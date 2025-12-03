package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	todoFile string
)

func init() {
	RootCmd.AddCommand(firstCmd)
	firstCmd.AddCommand(secondCmd)
	// RootCmd.AddCommand(addCmd)
}

var firstCmd = &cobra.Command{
	Use:   "first",
	Short: "fst",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You are at first command")
	},
}

var secondCmd = &cobra.Command{
	Use:   "second",
	Short: "snd",
	Run: func(cmd *cobra.Command, args []string) {
		color.Cyan("You are at second command")
	},
}
