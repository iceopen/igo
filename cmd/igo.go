package cmd

import (
	"igo/cmd/commands"
	_ "igo/cmd/commands/beego"
	_ "igo/cmd/commands/dep"
	_ "igo/cmd/commands/golang"
	_ "igo/cmd/commands/goimports"
	_ "igo/cmd/commands/version"
	"igo/utils"
)

var usageTemplate = `igo 帮你快速完成Golang开发助手`

var ErrorTemplate = `igo: %s.
用 {{"igo help" }} 或其他命令.
`

var helpTemplate = ``

func Usage() {
	utils.Tmpl(usageTemplate, commands.AvailableCommands)
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

func IfGenerateDocs(name string, args []string) bool {
	if name != "generate" {
		return false
	}
	for _, a := range args {
		if a == "docs" {
			return true
		}
	}
	return false
}
