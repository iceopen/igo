package goimports

import (
	"gitee.com/iceinto/igo/utils/command"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// 加载 goimports 工具

func NewGoimportsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "goimports",
		Short: "goimports 初始化",
		Run:   goimportsFn,
	}
	return cmd
}

// 执行核心
func goimportsFn(cmd *cobra.Command, args []string) {
	color.Red("goimports 工具相关，请先执行 golang")
	if len(args) == 1 && args[0] == "init" {
		// go get golang.org/x/tools/cmd/goimports
		PackageDownload()
	}
}

// 下载
func PackageDownload() {
	// 执行下载 go get -u github.com/golang/dep/cmd/dep
	color.Blue("下载 goimports 开始")
	command.Run("go", "get", "golang.org/x/tools/cmd/goimports")
	color.Blue("下载 goimports 结束")
}
