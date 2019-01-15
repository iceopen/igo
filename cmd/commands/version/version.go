package version

import (
	"os"
	"runtime"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// NewVersionCommand 获取当前版本号
func NewVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version",
		Run:   versionCommandFn,
	}

	return cmd
}

func versionCommandFn(cmd *cobra.Command, args []string) {
	runtimeInfo := RuntimeInfo{
		GetGoVersion(),
		runtime.GOOS,
		runtime.GOARCH,
		runtime.NumCPU(),
		os.Getenv("GOPATH"),
		runtime.GOROOT(),
		runtime.Compiler,
		version,
	}
	color.Red("", runtimeInfo)
}
