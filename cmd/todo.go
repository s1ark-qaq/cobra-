package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	todoFile string

	flag string
)

func init() {
	RootCmd.AddCommand(todoCmd)
	//varP 附带简写
	todoCmd.Flags().StringVarP(&flag, "flag", "f", "flag_default", "测试解析顺序")
	//var 不带简写
	todoCmd.Flags().StringVar(&flag, "flag", "flag_default", "测试解析顺序")
}

var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "创建todo文件",
	//显示args的数量为1，还有max和min
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var name = "todo"
		//args是没有-或--前缀的参数，支持灵活变动
		if len(args) > 0 {
			name = args[0]
		}
		todoFile = name + ".md"

		// 创建文件
		f, err := os.Create(todoFile)
		if err != nil {
			fmt.Println("创建文件失败:", err)
			return
		}
		defer f.Close()

		content := fmt.Sprintf("# %s list\n\n", name)
		_, _ = f.WriteString(content)
		fmt.Printf("%slist创建成功", name)
	},
}
