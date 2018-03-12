package beego

import (
	"igo/cmd/commands"

	"github.com/fatih/color"
)

// 加载 dep 工具

var CmdRun = &commands.Command{
	UsageLine: "bee [download] [init]",
	Short:     "bee",
	Long:      ``,
	PreRun:    func(cmd *commands.Command, args []string) {},
	Run:       RunApp,
}

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

// 执行核心
func RunApp(cmd *commands.Command, args []string) int {
	color.Red("bee 工具相关")
	if len(args) == 0 || args[0] == "download" {
		// 执行下载 go get github.com/astaxie/beego
		//go get github.com/beego/bee
	}

	return 0
}
