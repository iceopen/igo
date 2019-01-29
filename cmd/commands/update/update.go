package update

import (
	"github.com/fatih/color"
	"github.com/iceopen/igo/pkg/utils/command"
	"github.com/spf13/cobra"
)

// NewGolangCommand 加载 dep 工具
func NewUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "igo 升级",
		Run:   golangFn,
	}
	return cmd
}

// 执行核心
func golangFn(cmd *cobra.Command, args []string) {
	color.Red("igo 更新开始 ")
	command.Run("go", "get", "-u", "-v", "github.com/iceopen/igo")
	color.Red("igo 更新完成 ")
}
