package lemin

import (
	"fmt"
	"os"
)

// RunProgramWithFile - path is filepath
func RunProgramWithFile(path string) {
	fileContent, err := getFileContent(path)
	if err != nil {
		CloseProgram(err)
	}
	RunProgramWithContent(fileContent)
}

// GetResultByContent - Using for web.
func GetResultByContent(content string) (string, error) {
	lines := splitToLines(content)
	terrain, err := getTerrainFromLines(lines)
	if err != nil {
		return "", err
	}
	err = terrain.Validate()
	if err != nil {
		return "", err
	}
	// Add Start Match witch get result by lines or string
	PrintTerrainDatas(terrain)
	return "Correct", nil
}

// RunProgramWithContent - content is filecontent
func RunProgramWithContent(content string) {
	result, err := GetResultByContent(content)
	if err != nil {
		CloseProgram(err)
	}
	fmt.Println(result)
}

// CloseProgram - Closing program. And if `error == nil` exit code will be 0
func CloseProgram(err error) {
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
