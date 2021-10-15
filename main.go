package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"lem-in/lemin"
	"lem-in/web"
)

// Changes
// Completed Writing Result
// Fixed reading data for line
// Added Error handle for empty files and missing data on file/content
// Added comments for every func, struct
// The size of the footer is made of rubber (It was not beautiful on large screens)

// TODO Prevent memory leak for the web part
// TODO (HTML & README.md) Write About Project

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: Program takes only 1 argument (fileName or --flags)!")
		os.Exit(1)
	}
	argument := os.Args[1]
	//Set Flags
	port := flag.String("http", "", "--http=:port\n 1050 < port < 65000\n")
	filename := flag.String("file", "", "--file=filename\n")
	flag.Parse()

	//Start Program
	if strings.HasPrefix(argument, "--") { //if has flags
		if *port != "" {
			web.RunServer(*port)
		} else if *filename != "" {
			lemin.RunProgramWithFile(*filename)
		} else { //Default = Help
			flag.Usage()
			os.Exit(1)
		}
	} else { // Default
		lemin.RunProgramWithFile(argument)
	}
}
