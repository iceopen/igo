package main

import (
	"fmt"
	"os"

	"igo/cmd/commands/dep"
	"igo/cmd/commands/version"

	"igo/cmd/commands/beego"

	"igo/cmd/commands/goimports"

	"igo/cmd/commands/golang"

	"igo/cmd/commands/goinit"

	"github.com/spf13/cobra"
)

const (
	cliName        = "igo"
	cliDescription = "igo golang 开发辅助工具!"
)

var (
	rootCmd = &cobra.Command{
		Use:        cliName,
		Short:      cliDescription,
		SuggestFor: []string{"igo"},
	}
)

func init() {
	rootCmd.AddCommand(version.NewVersionCommand())
	rootCmd.AddCommand(dep.NewDepCommand())
	rootCmd.AddCommand(beego.NewBeeCommand())
	rootCmd.AddCommand(goimports.NewGoimportsCommand())
	rootCmd.AddCommand(golang.NewGolangCommand())
	rootCmd.AddCommand(goinit.NewInitCommand())
}

func main() {
	// TODO 添加环境依赖判断 golang 和 git 安装配置
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
