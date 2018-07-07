package goinit

import (
	"igo/cmd/commands/beego"
	"igo/cmd/commands/dep"
	"igo/cmd/commands/goimports"
	"igo/cmd/commands/golang"

	"github.com/spf13/cobra"
)

// 初始化  命令
func NewInitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "初始化 golang 开发环境基础依赖",
		Run:   initCommandFn,
	}
	return cmd
}

// 执行核心
func initCommandFn(cmd *cobra.Command, args []string) {
	beego.PackageDownload()
	dep.PackageDownload()
	golang.PackageDownload()
	goimports.PackageDownload()
}
