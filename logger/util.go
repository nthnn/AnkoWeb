package logger

import (
	"fmt"
	"strconv"
	"time"
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
