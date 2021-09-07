package main

import (
	"fmt"
	"os"

	"lem-in/lemin"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: Program takes only 1 argument (fileName)!")
		return
	}
	filePath := os.Args[1]
	lemin.RunProgramWithFile(filePath)
}
