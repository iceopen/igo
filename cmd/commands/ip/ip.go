package ip

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/spf13/cobra"
)

// NewIpCommand 加载 gops 工具
func NewIpCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ip",
		Short: "ip 获取当前网卡信息",
		Run:   getCmdNet,
	}
	return cmd
}

// 执行核心
func getCmdNet(cmd *cobra.Command, args []string) {
	getInterfaces()
}

func getInterfaces() {
	s, _ := net.Interfaces()
	if len(s) > 0 {
		for key, value := range s {
			color.Blue(" 网卡 %d 信息是：", key)
			fmt.Println(value)
		}
	} else {
		fmt.Println("获取错误")
	}
}
