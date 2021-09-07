package lemin

import (
	"fmt"
	"os"
	"strings"
)

// RunProgramWithFile - path is filepath
func RunProgramWithFile(path string) {
	fileContent, err := getFileContent(path)
	if err != nil {
		CloseProgram(err)
	}
	RunProgramWithContent(fileContent)
}

// RunProgramWithContent - content is filecontent
func RunProgramWithContent(content string) {
	lines := strings.Split(content, "\n")
	terrain, err := getTerrainFromLines(lines)
	fmt.Printf("Terrain: %q\nError: %v\n", terrain, err)
}

// CloseProgram - Closing program. And if `error == nil` exit code will be 0
func CloseProgram(err error) {
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
