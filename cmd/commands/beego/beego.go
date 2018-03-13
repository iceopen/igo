package beego

import (
	"igo/cmd/commands"

	"igo/utils/command"

	"github.com/fatih/color"
)

// 加载 beego 工具

var CmdRun = &commands.Command{
	UsageLine: "beego [download] [init]",
	Short:     "beego",
	Long:      ``,
	PreRun:    func(cmd *commands.Command, args []string) {},
	Run:       RunApp,
}

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

// 执行核心
func RunApp(cmd *commands.Command, args []string) int {
	color.Red("beego 工具相关")
	if len(args) == 1 && args[0] == "download" {
		// 执行下载 go get github.com/astaxie/beego
		//go get github.com/beego/bee
		PackageDownload()
	}
	return 0
}

// 下载
func PackageDownload() {
	// 执行下载 go get -u github.com/golang/dep/cmd/dep
	color.Blue("下载 beego 开始")
	command.Run("go", "get", "github.com/astaxie/beego")
	command.Run("go", "get", "github.com/beego/bee")
	color.Blue("下载 beego 结束")
}
