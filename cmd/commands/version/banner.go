package version

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"text/template"
	"time"
)

const version = "2.0.0"

// RuntimeInfo holds information about the current runtime.
type RuntimeInfo struct {
	GoVersion  string
	GOOS       string
	GOARCH     string
	NumCPU     int
	GOPATH     string
	GOROOT     string
	Compiler   string
	IgoVersion string
}

// InitBanner loads the banner and prints it to output
// All errors are ignored, the application will not
// print the banner in case of error.
func InitBanner(out io.Writer, in io.Reader) {
	if in == nil {
		log.Fatal("The input is nil")
	}

	banner, err := ioutil.ReadAll(in)
	if err != nil {
		log.Fatal(err)
	}

	show(out, string(banner))
}

func show(out io.Writer, content string) {
	t, err := template.New("banner").
		Funcs(template.FuncMap{"Now": Now}).
		Parse(content)

	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(out, RuntimeInfo{
		GetGoVersion(),
		runtime.GOOS,
		runtime.GOARCH,
		runtime.NumCPU(),
		os.Getenv("GOPATH"),
		runtime.GOROOT(),
		runtime.Compiler,
		version,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
}

// Now returns the current local time in the specified layout
func Now(layout string) string {
	return time.Now().Format(layout)
}

func GetGoVersion() string {
	var (
		cmdOut []byte
		err    error
	)

	if cmdOut, err = exec.Command("go", "version").Output(); err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(cmdOut), " ")[2]
}
