package dep

import (
	"igo/cmd/commands"

	"igo/utils/command"

	"strings"

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
	if len(args) == 1 {
		if args[0] == "download" {
			PackageDownload()
			return 0
		}
		//判断命令是否可用，如果可用提示直接使用dep命令。如果，不可以使用提示先进行下载操作
		info := command.Run("dep", "version")
		if strings.Contains(info, "dep:") {
			color.Blue("存在 dep ")
		} else {
			color.Red("请先执行，dep download")
		}
	}
	return 0
}

// 下载
func PackageDownload() {
	// 执行下载 go get -u github.com/golang/dep/cmd/dep
	color.Blue("下载 dep 开始")
	command.Run("go", "get", "-u", "github.com/golang/dep/cmd/dep")
	color.Blue("下载 dep 结束")
}
