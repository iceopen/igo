package vscode

import (
	"github.com/fatih/color"
	"github.com/iceopen/igo/pkg/utils/command"
	"github.com/spf13/cobra"
)

// VsCodeCommand 加载 VSCode 初始化
func VsCodeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vscode",
		Short: "vscode 环境初始化",
		Run:   getCmdNet,
	}
	return cmd
}

// 执行核心
//go get -v github.com/ramya-rao-a/go-outline
//go get -v github.com/acroca/go-symbols
//go get -v github.com/mdempsky/gocode
//go get -v github.com/rogpeppe/godef
//go get -v golang.org/x/tools/cmd/godoc
//go get -v github.com/zmb3/gogetdoc
//go get -v golang.org/x/lint/golint
//go get -v github.com/fatih/gomodifytags
//go get -v golang.org/x/tools/cmd/gorename
//go get -v sourcegraph.com/sqs/goreturns
//go get -v golang.org/x/tools/cmd/goimports
//go get -v github.com/cweill/gotests/...
//go get -v golang.org/x/tools/cmd/guru
//go get -v github.com/josharian/impl
//go get -v github.com/haya14busa/goplay/cmd/goplay
//go get -v github.com/uudashr/gopkgs/cmd/gopkgs
//go get -v github.com/davidrjenni/reftools/cmd/fillstruct
func getCmdNet(cmd *cobra.Command, args []string) {
	color.Blue("环境初始化中...")
	command.Run("go", "get", "github.com/ramya-rao-a/go-outline")
	command.Run("go", "get", "github.com/acroca/go-symbols")
	command.Run("go", "get", "github.com/mdempsky/gocode")
	command.Run("go", "get", "github.com/rogpeppe/godef")
	command.Run("go", "get", "github.com/zmb3/gogetdoc")
	command.Run("go", "get", "golang.org/x/lint/golint")
	command.Run("go", "get", "github.com/fatih/gomodifytags")
	command.Run("go", "get", "golang.org/x/tools/cmd/gorename")
	command.Run("go", "get", "github.com/sqs/goreturns")
	command.Run("go", "get", "golang.org/x/tools/cmd/goimports")
	command.Run("go", "get", "github.com/cweill/gotests/...")
	command.Run("go", "get", "golang.org/x/tools/cmd/guru")
	command.Run("go", "get", "github.com/josharian/impl")
	command.Run("go", "get", "github.com/haya14busa/goplay/cmd/goplay")
	command.Run("go", "get", "github.com/uudashr/gopkgs/cmd/gopkgs")
	command.Run("go", "get", "github.com/davidrjenni/reftools/cmd/fillstruct")
	color.Red("环境初始化完成")
}
