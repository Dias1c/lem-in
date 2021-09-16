package web

import (
	"fmt"
	"os"
)

// RunServer - starts server with setted port
func RunServer(port string) {
	err := validatePort(port)
	if err != nil {
		CloseProgram(err)
	}
	// Init templates + Handlers + Run Server (To Do)
}

// CloseProgram - Closing program. And if `error == nil` exit code will be 0
func CloseProgram(err error) {
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
