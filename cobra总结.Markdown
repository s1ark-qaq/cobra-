对cobra的自我总计：
cobra是基于go的spf13/pflag库的封装，有命令树（根命令+子命令的形式），xxxCli-help（自动生成-help命令），可以使用
cobra包中doc一键生成命令注释文档。其中的pflag支持string、bool、int类型等的传参，PersistentFlags还支持自动继承给
子命令，当子命令的flag跟继承的flag冲突时，自动选择当前命令绑定flag。
搭配gin框架可以一键启动web服务，并且初始化各种实例传入router中。

优点：
1.可以设置子命令，以及子命令的子命令。
2.可以给命令设置flag，丰富命令输入时的参数。例如：在项目初始化时指定配置文件。
3。自动生成--help命令获取命令树以及简短说明，有利于快速上手，也可以在代码中使用 cmd.Help()来手动打印帮助信息。
