package echo

import (
	"github.com/fatih/color"
	"github.com/iceopen/igo/pkg/utils/command"
	"github.com/spf13/cobra"
)

// NewEchoCommand 加载 echo 工具
func NewEchoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "echo",
		Short: "echo 初始化",
		Run:   echoCommandFn,
	}
	return cmd
}

// 执行核心
func echoCommandFn(cmd *cobra.Command, args []string) {
	color.Red("echo 工具相关")
	if len(args) == 1 && args[0] == "init" {
		PackageDownload()
	}
}

// PackageDownload 下载
func PackageDownload() {
	// 执行下载 go get -u github.com/golang/dep/cmd/dep
	color.Blue("下载 echo 开始")
	command.Run("go", "get", "github.com/labstack/echo")
	command.Run("go", "get", "github.com/dgrijalva/jwt-go")
	color.Blue("下载 echo 结束")
}
