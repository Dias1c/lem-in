package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Dias1c/lem-in/internal/lemin"
	"github.com/Dias1c/lem-in/internal/web"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: Program takes argument (fileName or --flags)!")
		os.Exit(1)
	}
	argument := os.Args[1]
	// Set Flags
	port := flag.String("http", "", "--http=:port\n 1050 < port < 65000\n")
	filename := flag.String("file", "", "--file=filename\n")
	flag.Parse()

	// Start Program
	if strings.HasPrefix(argument, "--") { // if has flags
		if *port != "" {
			err := web.RunServer(*port)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
		} else if *filename != "" {
			lemin.RunProgramWithFile(*filename, true)
		} else { // Default = Help
			flag.Usage()
			os.Exit(1)
		}
	} else { // Default
		lemin.RunProgramWithFile(argument, false)
	}
}
