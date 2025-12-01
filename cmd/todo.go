package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	todoFile string
)

func init() {
	RootCmd.AddCommand(todoCmd)
	todoCmd.AddCommand(addCmd)
	todoCmd.AddCommand(downCmd)
	todoCmd.AddCommand(listCmd)
}

var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "创建todo文件",
	Run: func(cmd *cobra.Command, args []string) {
		var name = "todo"
		if len(args) > 0 {
			name = args[0]
		}
		todoFile = name + ".md"

		// 检查文件是否已存在
		if _, err := os.Stat(todoFile); err == nil {
			fmt.Println("文件已存在:", todoFile)
			return
		}

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

var addCmd = &cobra.Command{
	Use:   "new",
	Short: "新增todo",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := strings.Join(args, " ")

		if todoFile == "" {
			fmt.Println("未创建todo文件，请先创建")
			return
		}

		f, err := os.OpenFile(todoFile, os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		_, _ = f.WriteString(fmt.Sprintf("- [ ] %s\n", title))
	},
}

var downCmd = &cobra.Command{
	Use:   "d",
	Short: "完成任务",
	Run: func(cmd *cobra.Command, args []string) {
		if todoFile == "" {
			fmt.Println("未创建todo文件，请先创建")
			return
		}

		data, err := os.ReadFile(todoFile)
		if err != nil {
			fmt.Println(err)
			return
		}

		lines := strings.Split(string(data), "\n")
		num := -1
		fmt.Sscanf(args[0], "%d", &num)
		if num <= 0 || num > len(lines) {
			fmt.Println("无效的任务编号")
			return
		}

		// 查找第 num 个任务行
		count := 0
		for i, line := range lines {
			if strings.HasPrefix(line, "- [ ]") || strings.HasPrefix(line, "- [x]") {
				count++
				if count == num {
					lines[i] = strings.Replace(line, "- [ ]", "- [x]", 1)
					break
				}
			}
		}

		err = os.WriteFile(todoFile, []byte(strings.Join(lines, "\n")), 0644)
		if err != nil {
			fmt.Println("写入文件失败:", err)
			return
		}

		fmt.Println("任务完成:", args[0])
	},
}

var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "todo列表",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Open(todoFile)
		if err != nil {
			fmt.Println("无法打开文件:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		taskNum := 0
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "- [") {
				taskNum++
				fmt.Printf("%d. %s\n", taskNum, line)
			} else {
				fmt.Println(line)
			}
		}
	},
}
