package main

import (
	"fmt"
	"github.com/iceopen/igo/cmd/commands/vscode"
	"os"

	"github.com/iceopen/igo/cmd/commands/beego"
	"github.com/iceopen/igo/cmd/commands/echo"
	"github.com/iceopen/igo/cmd/commands/goimports"
	"github.com/iceopen/igo/cmd/commands/goinit"
	"github.com/iceopen/igo/cmd/commands/gops"
	"github.com/iceopen/igo/cmd/commands/ip"
	"github.com/iceopen/igo/cmd/commands/update"
	"github.com/iceopen/igo/cmd/commands/version"
	"github.com/spf13/cobra"
)

const (
	cliName        = "igo"
	cliDescription = "igo Go 开发环境辅助工具!"
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
	rootCmd.AddCommand(beego.NewBeeCommand())
	rootCmd.AddCommand(goimports.NewGoimportsCommand())
	rootCmd.AddCommand(goinit.NewInitCommand())
	rootCmd.AddCommand(echo.NewEchoCommand())
	rootCmd.AddCommand(gops.NewGoPsCommand())
	rootCmd.AddCommand(ip.NewIpCommand())
	rootCmd.AddCommand(update.NewUpdateCommand())
	rootCmd.AddCommand(vscode.VsCodeCommand())
}

func main() {
	// TODO 添加环境依赖判断 golang 和 git 安装配置
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
