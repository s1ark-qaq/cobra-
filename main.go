package main

import (
	"bufio"
	"cobra/cmd"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Todo启动, 输入 'exit' 退出")

	for {
		if !scanner.Scan() {
			break // Ctrl+D
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}
		if input == "exit" {
			break
		}

		// 将用户输入作为命令行参数传给 Cobra
		args := strings.Fields(input)
		cmd.RootCmd.SetArgs(args)
		cmd.Execute()
	}

}
