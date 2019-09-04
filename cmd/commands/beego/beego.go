package beego

import (
	"github.com/fatih/color"
	"github.com/iceopen/igo/pkg/utils/command"
	"github.com/spf13/cobra"
)

// NewBeeCommand 加载 beego 工具
func NewBeeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "beego",
		Short: "beego 初始化",
		Run:   beegoCommandFn,
	}
	return cmd
}

// beegoCommandFn 执行核心
func beegoCommandFn(cmd *cobra.Command, args []string) {
	color.Red("beego 工具相关")
	if len(args) == 1 && args[0] == "init" {
		PackageDownload()
	}
}

// PackageDownload 下载
func PackageDownload() {
	color.Blue("下载 beego 开始")
	command.Run("go", "get", "github.com/astaxie/beego")
	command.Run("go", "get", "github.com/beego/bee")
	color.Blue("下载 beego 结束")
}
