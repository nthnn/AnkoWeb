package logger

import (
	"github.com/fatih/color"
)

func Plain(message string) {
	c := color.New(color.FgGreen)
	c.Print("[" + timeNow() + "]")

	printMessage(message)
}

func Error(message string) {
	c := color.New(color.FgRed)
	c.Print("[" + timeNow() + "]")

	printMessage(message)
}
