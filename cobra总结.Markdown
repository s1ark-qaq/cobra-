# 对cobra的自我总计
cobra是基于go的spf13/pflag库的封装，用于快速构建cli应用。有命令树（根命令+子命令的形式），xxxCli-help（自动生成-help命令），可以使用
cobra包中doc一键生成命令注释文档。其中的pflag支持string、bool、int类型等的传参，PersistentFlags还支持自动继承给
子命令，当子命令的flag跟继承的flag冲突时，自动选择当前命令绑定flag。
搭配gin框架可以一键启动web服务，并且初始化各种实例传入router中。

## 大概运行流程
在程序开始时init函数中执行addcmd构建命令树，cmd.AddCmd(subCmd)，subCmd加入cmd的commands内部，subCmd.parent = cmd。

rootCmd.Execute是整个命令行的入口，使用go标准库os.Args获取用户的输入，取args[1:]的全部参数，通过c.Find
对args进行处理，使用innerfind递归处理args，找到最终的command和剩余的args。执行c.execute

### rootCmd.Execute
这一步会递归找到执行主体的父命令最终执行根命令的execute

### c.Find
find中对rootCmd.Commands遍历，args中字符串和commands.name匹配则返回subCommands。

### c.execute
最终执行主体，各种run有先后顺序，设置了RunE则Run会被忽略。
执行顺序为c.preRun()-->c.Run()-->c.postRun()

### cmd.AddCmd(subCmd)
将输入subCmd加入cmd.Commands([]*Command)，在执行函数时遍历commands拿到被执行命令。

## 优点
1.可以设置子命令，以及子命令的子命令。使用c.Commands和parent来管理命令。
2.可以给命令设置flag，丰富命令输入时的参数。例如：在项目初始化时指定配置文件。
3。自动生成--help命令获取命令树以及简短说明，有利于快速上手，也可以在代码中使用 cmd.Help()来手动打印帮助信息。

## 使用cobra的大型项目
doker,githubCli,Hugo等
这些项目选择cobra的主要原因:
1.项目本身使用go语言开发,而cobra又是go语言用于构建go cli的库
2.这些项目拥有许多层级的子命令,而cobra的命令树易于管理
例如github cli--https://cli.github.com/manual/?utm_source=chatgpt.com  拥有丰富的一级和二级子命令
3.自动生成help命令,有助于大型项目的维护