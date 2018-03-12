package dep

import (
	"igo/cmd/commands"

	"igo/utils/command"

	"github.com/fatih/color"
)

// 加载 dep 工具

var CmdRun = &commands.Command{
	UsageLine: "dep [download] [init]",
	Short:     "dep",
	Long:      ``,
	PreRun:    func(cmd *commands.Command, args []string) {},
	Run:       RunApp,
}

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

// 执行核心
func RunApp(cmd *commands.Command, args []string) int {
	color.Red("dep 工具相关")
	if args[0] == "download" {
		DepDownload()
	}
	return 0
}

// 下载
func DepDownload() {
	// 执行下载 go get -u github.com/golang/dep/cmd/dep
	color.Blue("下载 dep 开始")
	command.Run("go", "get", "-u", "github.com/golang/dep/cmd/dep")
	color.Blue("下载 dep 结束")
}
