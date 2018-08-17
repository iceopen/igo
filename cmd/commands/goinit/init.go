package goinit

import (
	"gitee.com/iceinto/igo/cmd/commands/beego"
	"gitee.com/iceinto/igo/cmd/commands/goimports"
	"gitee.com/iceinto/igo/cmd/commands/golang"
	"github.com/spf13/cobra"
	"gitee.com/iceinto/igo/cmd/commands/dep"
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
	golang.PackageDownload()
	beego.PackageDownload()
	goimports.PackageDownload()
	dep.PackageDownload()
}
