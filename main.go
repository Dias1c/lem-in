package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"lem-in/lemin"
)

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
			fmt.Printf("Work Port in %q\nNot working now!\n", *port)
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
