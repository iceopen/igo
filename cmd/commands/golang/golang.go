package golang

import (
	"os"

	"github.com/fatih/color"
	"github.com/iceopen/igo/pkg/utils"
	"github.com/iceopen/igo/pkg/utils/command"
	"github.com/spf13/cobra"
)

// NewGolangCommand 加载 dep 工具
func NewGolangCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "golang",
		Short: "golang 初始化",
		Run:   golangFn,
	}
	return cmd
}

// 执行核心
func golangFn(cmd *cobra.Command, args []string) {
	color.Red("golang 环境初始化 ")
	if len(args) == 1 {
		if args[0] == "init" {
			PackageDownload(false)
		} else if args[0] == "update" {
			PackageDownload(true)
		} else {
			color.Red("请先执行，golang init")
		}
	}
}

// PackageDownload 下载
func PackageDownload(update bool) {
	xPath := utils.GetGOPATHs()[0] + "/src/golang.org/x"
	// 判断目录是否存在
	isExist := utils.IsExist(xPath)
	if isExist == false {
		// 创建目录
		utils.CreateDirs(xPath)
	}
	// https://github.com/golang/net.git
	strs := []string{"net", "tools", "sys", "crypto", "text", "image", "mock", "snappy", "lint"}
	for _, v := range strs {
		if (utils.IsExist(v) == false && utils.IsExist(xPath+"/"+v) == false) || update {
			color.Blue("下载 " + v + " 开始")
			os.RemoveAll("./" + v)
			command.Run("git", "clone", "https://github.com/golang/"+v+".git")
			color.Blue("下载 " + v + " 结束")
		}
		if utils.IsExist(xPath+"/"+v) == false || update {
			color.Blue("移动 " + v + " 目录")
			// 移除老的文件目录
			os.RemoveAll(xPath + "/" + v)
			os.Rename(v, xPath+"/"+v)
			color.Blue("移动 " + v + " 结束")
		}
		// 删除本地缓存文件目录
		os.RemoveAll("./" + v)
	}
}
