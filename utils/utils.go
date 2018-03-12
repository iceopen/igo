package utils

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

func Tmpl(text string, data interface{}) {
	t := template.New("Usage")
	template.Must(t.Parse(text))
	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func PrintErrorAndExit(message, errorTemplate string) {
	Tmpl(fmt.Sprintf(errorTemplate, message), nil)
	os.Exit(2)
}
