package process

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/process"
)

// 获取进程列表
func getProcessList() {
	p, _ := process.Processes()
	if len(p) > 0 {
		for _, value := range p {
			pName, _ := value.Name()
			pUserName, _ := value.Username()
			pNumThreads, _ := value.NumThreads()
			color.Blue("进程ID：%d，进程名称：%s，用户：%s %d", value.Pid, pName, pUserName, pNumThreads)
			fmt.Println(value.MemoryInfo())
		}
	} else {
		fmt.Println("获取错误")
	}
}

func getProcessPids() {
	p, _ := process.Pids()
	fmt.Println(p)
}
