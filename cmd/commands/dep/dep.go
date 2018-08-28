package dep

import (
	"strings"

	"gitee.com/iceinto/igo/utils/command"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// 初始化 dep 命令
func NewDepCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dep",
		Short: "dep 初始化",
		Run:   depCommandFn,
	}
	return cmd
}

// 执行核心
func depCommandFn(cmd *cobra.Command, args []string) {
	color.Red("dep 工具相关")
	if len(args) == 1 {
		if args[0] == "init" {
			PackageDownload()
		}
		// 判断命令是否可用，如果可用提示直接使用dep命令。如果，不可以使用提示先进行下载操作
		info := command.Run("dep", "version")
		if strings.Contains(info, "dep:") {
			color.Blue("存在 dep ")
		} else {
			color.Red("请先执行，dep init")
		}
	}
}

// 下载
func PackageDownload() {
	// 执行下载 go get -u github.com/golang/dep/cmd/dep
	color.Blue("下载 dep 开始")
	command.Run("go", "get", "-v", "github.com/golang/dep/cmd/dep")
	color.Blue("下载 dep 结束")
}
