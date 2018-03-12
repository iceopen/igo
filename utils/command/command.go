package command

import (
	"log"
	"os/exec"
)

// 运行命令行
func Run(name string, arg ...string) string {
	cmd := exec.Command(name, arg...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}
