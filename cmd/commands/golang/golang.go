package dep

import (
	"igo/cmd/commands"
	"igo/utils/command"
	"github.com/fatih/color"
	"igo/utils"
)

// 加载 dep 工具

var CmdRun = &commands.Command{
	UsageLine: "golang [init]",
	Short:     "golang",
	Long:      ``,
	PreRun:    func(cmd *commands.Command, args []string) {},
	Run:       RunApp,
}

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

// 执行核心
func RunApp(cmd *commands.Command, args []string) int {
	color.Red("golang 工具相关 ")
	if len(args) == 1 {
		if args[0] == "init" {
			PackageDownload()
			return 0
		}
		color.Red("请先执行，dep init")
	}
	return 0
}

// 下载
func PackageDownload() {
	xPath := utils.GetGOPATHs()[0] + "/src/golang.org/x"
	// 判断目录是否存在
	isExist := utils.IsExist(xPath)
	if isExist == false {
		// 创建目录
		utils.CreateDirs(xPath)
	}
	// https://github.com/golang/net.git
	// 执行下载 go get -u github.com/golang/dep/cmd/dep
	strs := []string{"net", "tools", "sys", "crypto", "text"}
	for _, v := range strs {
		if utils.IsExist(v) == false && utils.IsExist(xPath+"/"+v) == false {
			color.Blue("下载 " + v + " 开始")
			command.Run("git", "clone", "https://github.com/golang/"+v+".git")
			color.Blue("下载 " + v + " 结束")
		}
		if utils.IsExist(xPath+"/"+v) == false {
			color.Blue("移动 " + v + " 目录")
			command.Run("mv", v, xPath)
			color.Blue("移动 " + v + " 结束")
		}
	}

}
