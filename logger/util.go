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

func colorize(fgColor color.Attribute) func(a ...interface{}) (int, error) {
	c := color.New(fgColor)
	return c.Print
}

func printMessage(message string) {
	fmt.Print(" ")
	fmt.Println(message)
}
