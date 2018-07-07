package golang

import (
	"igo/utils"
	"igo/utils/command"

	"runtime"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// 加载 dep 工具
func NewGolangCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "goimports",
		Short: "goimports 初始化",
		Run:   golangFn,
	}
	return cmd
}

// 执行核心
func golangFn(cmd *cobra.Command, args []string) {
	color.Red("golang 工具相关 ")
	if len(args) == 1 {
		if args[0] == "init" {
			PackageDownload()
		}
		color.Red("请先执行，golang init")
	}
}

// 下载
func PackageDownload() {
	mvStr := "mv"
	if runtime.GOOS == "windows" {
		mvStr = "move"
	}
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
			command.Run(mvStr, v, xPath)
			color.Blue("移动 " + v + " 结束")
		}
	}

}
