package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/JesseyBland/numbers/api"
	"github.com/JesseyBland/numbers/app"
	"github.com/JesseyBland/numbers/html"
)

func main() {

	CharPTR := flag.String("char", "#", "character returned")
	PortPTR := flag.String("port", "9000", "port")
	// Set up error log to ouptut information
	log.SetFlags(log.Lshortfile)

	// Parse Command Line arguements and use them to find
	// appropriate sub-command for cnum to run.
	flag.Parse()
	switch flag.Arg(0) {
	case "api":
		api.Server(app.SetPort(*PortPTR), *CharPTR)

	case "html":
		html.Server(app.SetPort(*PortPTR), *CharPTR)

	case "local":
		cnum := app.Local(flag.Arg(1), *CharPTR)
		fmt.Println(cnum)
	case "help":
		fmt.Println(usage)
	default:
		fmt.Println("No valid sub command selected. Use \"cnum help\" for information.")
	}
}
