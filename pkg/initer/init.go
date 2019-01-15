package initer

import (
	"log"
	"runtime"

	"github.com/google/gops/agent"
)

// init 初始化项目应用需要做的内容
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// GoPsAgent gops is a command to list and diagnose Go processes currently running on your system.
func GoPsAgent() {
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	} else {
		log.Println("gops/agent start")
	}
}
