package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version string

var RootCmd = &cobra.Command{
	Use:     "gogo",  //根命令
	Short:   "gogo",  //简短解释
	Version: version, //支持gogo version获取版本信息
	Run: func(cmd *cobra.Command, args []string) { //命令执行主体
		//无前缀参数被解析成字符串数组
		fmt.Println("gogo命令参数:", args[0])
	},
}

func init() {
	//绑定flag到指针，v是简写，value是默认值，最后是介绍
	RootCmd.PersistentFlags().StringVarP(&version, "version", "v", "v0.0.1", "set_version")
}

func Execute() {
	if err := RootCmd.Execute(); err != nil { //根命令运行
		fmt.Println(err)
		os.Exit(1)
	}
}
