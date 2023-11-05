package logger

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fatih/color"
)

func timeNow() string {
	now := ""
	current := time.Now()

	now += strconv.Itoa(current.Year()) + " "
	now += current.Month().String() + " "
	now += strconv.Itoa(current.Day()) + " - "

	now += strconv.Itoa(current.Hour()) + ":"
	now += strconv.Itoa(current.Minute()) + ":"
	now += strconv.Itoa(current.Second()) + "."
	now += strconv.Itoa(int(current.UnixMilli()))[0:3]

	return now
}

func printMessage(message string) {
	fmt.Print(" ")
	fmt.Println(message)
}

func Plain(message string) {
	c := color.New(color.FgGreen)
	c.Print("[" + timeNow() + "]")

	printMessage(message)
}

func Error(message string) {
	c := color.New(color.FgRed)
	c.Print(time.Now().String())

	printMessage(message)
}
