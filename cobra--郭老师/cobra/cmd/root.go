package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string

var RootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) { //命令执行主体
		fmt.Println("This is root command")
	},
}

func init() {
	// RootCmd.PersistentFlags().StringVarP(&version, "version", "v", "v0.0.1", "set_version")
}

func Execute() {
	RootCmd.Execute()
}
