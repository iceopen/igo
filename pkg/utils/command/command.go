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
		log.Println("命令行执行错误")
		log.Fatal(err)
	}
	return string(out)
}
