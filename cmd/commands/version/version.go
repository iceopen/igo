package version

import (
	"fmt"
	"igo/cmd/commands"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"encoding/json"
	"log"
	"bytes"
	"gopkg.in/yaml.v2"
)

const verboseVersionBanner string = `%s%s______
igo v{{ .BeeVersion }}%s
%s%s
├── GoVersion : {{ .GoVersion }}
├── GOOS      : {{ .GOOS }}
├── GOARCH    : {{ .GOARCH }}
├── NumCPU    : {{ .NumCPU }}
├── GOPATH    : {{ .GOPATH }}
├── GOROOT    : {{ .GOROOT }}
├── Compiler  : {{ .Compiler }}
└── Date      : {{ Now "Monday, 2 Jan 2006" }}%s
`

var outputFormat string

const version = "1.0.0"

func init() {

}

func versionCmd(cmd *commands.Command, args []string) int {

	cmd.Flag.Parse(args)
	stdout := cmd.Out()

	if outputFormat != "" {
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
		switch outputFormat {
		case "json":
			{
				b, err := json.MarshalIndent(runtimeInfo, "", "    ")
				if err != nil {
					log.Fatal(err.Error())
				}
				fmt.Println(string(b))
				return 0
			}
		case "yaml":
			{
				b, err := yaml.Marshal(&runtimeInfo)
				if err != nil {
					log.Fatal(err.Error())
				}
				fmt.Println(string(b))
				return 0
			}
		}
	}

	coloredBanner := fmt.Sprintf(verboseVersionBanner, "\x1b[35m", "\x1b[1m",
		"\x1b[0m", "\x1b[32m", "\x1b[1m", "\x1b[0m")
	InitBanner(stdout, bytes.NewBufferString(coloredBanner))
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