package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version string

var rootCmd = &cobra.Command{
	Use:     "gogo",  //根命令
	Short:   "gogo",  //简短解释
	Version: version, //支持-v，-V，-version获取版本信息
	Run: func(cmd *cobra.Command, args []string) { //命令执行主体
		fmt.Println("gogo:", version)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&version, "version", "v", "v0.0.1", "set_version")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil { //根命令运行
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//在命令行运行之前，init之后的初始化。
	//在函数内部可以插入配置文件初始化逻辑，在程序启动之初配合命令设置或修改配置文件。
	cobra.OnInitialize()
}
