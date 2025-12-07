# go mod解读
go mod的相关命令,是简化的cobra版本主要使用command结构体来表示一个命令,其中commands属性表示子命令,run函数表示命令的执行主体.
在设计上和cobra类似,使用commands管理子命令


go mod 的cli管理:

type Command struct {

执行主体
Run func(ctx context.Context, cmd *Command, args []string)

命令使用名
UsageLine string

Short string

Long string

标志	
Flag flag.FlagSet

CustomFlags bool

子命令
Commands []*Command

}

## go mod init
go mod管理器的初始化,创建go.mod

## go mod tidy
这是go语言本身用于包管理的一个命令,也是go源码base.Command的子命令之一,作用是整理 go.mod 和 go.sum 文件
主要通过modload.LoadPackages函数进行包的扫描与下载缺失的包