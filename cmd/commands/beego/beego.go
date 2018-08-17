package beego

import (
	"github.com/fatih/color"
	"gitee.com/iceinto/igo/utils/command"
	"github.com/spf13/cobra"
)

// 加载 beego 工具

func NewBeeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "beego",
		Short: "beego 初始化",
		Run:   beegoCommandFn,
	}
	return cmd
}

// 执行核心
func beegoCommandFn(cmd *cobra.Command, args []string) {
	color.Red("beego 工具相关")
	if len(args) == 1 && args[0] == "init" {
		PackageDownload()
	}
}

// 下载
func PackageDownload() {
	// 执行下载 go get -u github.com/golang/dep/cmd/dep
	color.Blue("下载 beego 开始")
	command.Run("go", "get", "github.com/astaxie/beego")
	command.Run("go", "get", "github.com/beego/bee")
	color.Blue("下载 beego 结束")
}
