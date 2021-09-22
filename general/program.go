package general

import (
	"fmt"
	"os"
)

// CloseProgram - Closing program. And if `error == nil` exit code will be 0
func CloseProgram(err error) {
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
