package gops

import (
	"github.com/fatih/color"
	"github.com/iceopen/igo/pkg/utils/command"
	"github.com/spf13/cobra"
)

// NewGoPsCommand 加载 gops 工具
func NewGoPsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gops",
		Short: "gops 初始化",
		Run:   gopsFn,
	}
	return cmd
}

// 执行核心
func gopsFn(cmd *cobra.Command, args []string) {
	color.Red("gops init 使用")
	if len(args) == 1 && args[0] == "init" {
		// go get golang.org/x/tools/cmd/goimports
		PackageDownload()
	}
}

// PackageDownload 下载
func PackageDownload() {
	color.Blue("下载 gops 开始")
	command.Run("go", "install", "github.com/google/gops")
	color.Blue("下载 gops 结束")
}
