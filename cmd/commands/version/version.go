package version

import (
	"igo/cmd/commands"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/fatih/color"
)

var outputFormat string

const version = "1.0.0"

var CmdVersion = &commands.Command{
	UsageLine: "version",
	Short:     "Prints the current Bee version",
	Long: `
Prints the current Bee, Beego and Go version alongside the platform information.
`,
	Run: VersionCmd,
}

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdVersion)
}

func VersionCmd(cmd *commands.Command, args []string) int {
	color.Blue("版本和环境变量配置")

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
	return 0
}

func GetGoVersion() string {
	var (
		cmdOut []byte
		err    error
	)

	if cmdOut, err = exec.Command("go", "version").Output(); err != nil {
		log.Fatal("There was an error running 'go version' command: %s", err)
	}
	return strings.Split(string(cmdOut), " ")[2]
}
