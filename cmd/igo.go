package cmd

import (
	"igo/cmd/commands"
	"igo/utils"
)

var usageTemplate = `igo 帮你快速完成Golang开发助手`

var ErrorTemplate = `igo: %s.
用 {{"igo help" | bold}} 或其他命令.
`

var helpTemplate = ``

// 默认用户体系
func Usage() {
	utils.Tmpl(usageTemplate, "")
}

func Help(args []string) {
	if len(args) == 0 {
		Usage()
	}
	if len(args) != 1 {
		utils.PrintErrorAndExit("Too many arguments", ErrorTemplate)
	}

	arg := args[0]

	for _, cmd := range commands.AvailableCommands {
		if cmd.Name() == arg {
			utils.Tmpl(helpTemplate, cmd)
			return
		}
	}
	utils.PrintErrorAndExit("Unknown help topic", ErrorTemplate)
}
