package main

import (
	"cobra/cmd"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	fmt.Println("init")
}

func cobraInit() {
	fmt.Println("cobraInit")
}

func main() {
	//在命令行运行之前，init之后的初始化。
	//在函数内部可以插入配置文件初始化逻辑，在程序启动之初配合命令设置或修改配置文件。
	cobra.OnInitialize(cobraInit)
	//启动根命令
	cmd.Execute()
}
