package logger

import (
	"github.com/fatih/color"
)

func Plain(message string) {
	colorize(color.FgGreen)("[" + timeNow() + "]")
	printMessage(message)
}

func Error(message string) {
	colorize(color.FgRed)("[" + timeNow() + "]")
	printMessage(message)
}

func Info(message string) {
	colorize(color.FgCyan)("[" + timeNow() + "]")
	printMessage(message)
}
