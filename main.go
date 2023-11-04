package main

import (
	"flag"
	"fmt"
	"os"
)

func initFlags() {
	flag.CommandLine.SetOutput(os.Stdout)

	flag.Usage = func() {
		fmt.Println("AnkoWeb v0.0.1 - Primitive Version\n\nParameters:")
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	path := flag.String("path", ".", "Working directory for the server.")
	host := flag.String("host", "", "Name of the localhost server.")
	port := flag.Int("port", 1234, "Port of the localhost server.")
	initFlags()

	awpServer(*path, *host, int16(*port))
}
