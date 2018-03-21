package goimports

import (
	"igo/cmd/commands"
	"igo/utils/command"

	"github.com/fatih/color"
)

// 加载 goimports 工具

var CmdRun = &commands.Command{
	UsageLine: "goimports [download] [init]",
	Short:     "goimports",
	Long:      ``,
	PreRun:    func(cmd *commands.Command, args []string) {},
	Run:       RunApp,
}

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

// 执行核心
func RunApp(cmd *commands.Command, args []string) int {
	color.Red("goimports 工具相关")
	if len(args) == 1 && args[0] == "init" {
		// go get golang.org/x/tools/cmd/goimports
		PackageDownload()
	}
	return 0
}

// 下载
func PackageDownload() {
	// 执行下载 go get -u github.com/golang/dep/cmd/dep
	color.Blue("下载 goimports 开始")
	command.Run("go", "get", "golang.org/x/tools/cmd/goimports")
	color.Blue("下载 goimports 结束")
}
