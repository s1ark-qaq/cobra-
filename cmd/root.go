package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "v0.0.1"

var RootCmd = &cobra.Command{
	Use:     "gogo",  //根命令
	Short:   "gogo",  //简短解释
	Version: version, //支持gogo --version||-v 打印version内容
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root cmd")
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil { //根命令运行
		fmt.Println(err)
		os.Exit(1)
	}
}
