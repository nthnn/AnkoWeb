package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/nthnn/AnkoWeb/server"
)

func printBanner() {
	color.New(color.FgHiBlue).Print(`
█████╗ ███╗   ██╗██╗  ██╗ ██████╗ ██╗    ██╗███████╗██████╗ 
██╔══██╗████╗  ██║██║ ██╔╝██╔═══██╗██║    ██║██╔════╝██╔══██╗
███████║██╔██╗ ██║█████╔╝ ██║   ██║██║ █╗ ██║█████╗  ██████╔╝
██╔══██║██║╚██╗██║██╔═██╗ ██║   ██║██║███╗██║██╔══╝  ██╔══██╗
██║  ██║██║ ╚████║██║  ██╗╚██████╔╝╚███╔███╔╝███████╗██████╔╝
╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝ ╚═════╝  ╚══╝╚══╝ ╚══════╝╚═════╝ 
                                                             
`,
	)
}

func initFlags() {
	flag.CommandLine.SetOutput(os.Stdout)

	flag.Usage = func() {
		printBanner()

		color.New(color.FgHiYellow).Print("              AnkoWeb v0.0.1 - Primitive Version\n\n")

		fmt.Println("–––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––")
		fmt.Println("Parameters:")
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	path := flag.String("path", ".", "Working directory for the server.")
	host := flag.String("host", "", "Name of the localhost server.")
	port := flag.Int("port", 1234, "Port of the localhost server.")

	initFlags()
	server.AwpServer(*path, *host, int16(*port))
}
