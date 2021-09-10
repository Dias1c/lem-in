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
	if err != nil {
		CloseProgram(err)
	}
	fmt.Println("Rooms: Name\tX:\tY:")
	for _, room := range terrain.Rooms {
		fmt.Printf("Room: %v  \t%d\t%d\n", room.Name, room.X, room.Y)
	}
	fmt.Printf("StartRoom: %v\nEndRoom: %v\nCountAnts: %v\nError: %v\n", terrain.Start, terrain.End, terrain.AntsCount, err)
	// fmt.Printf("Terrain: %q\nError: %v\n", terrain, err)
}

// CloseProgram - Closing program. And if `error == nil` exit code will be 0
func CloseProgram(err error) {
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
