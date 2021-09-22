package lemin

import (
	"fmt"

	"lem-in/general"
)

// RunProgramWithFile - path is filepath
func RunProgramWithFile(path string) {
	fileContent, err := getFileContent(path)
	if err != nil {
		general.CloseProgram(err)
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
	result, err := terrain.Match()
	if err != nil {
		return "Incorrect", err
	}
	fmt.Println(result)

	return fmt.Sprintf("%v\n\n#Result:\n%v", content, result), nil
}

// RunProgramWithContent - content is filecontent
func RunProgramWithContent(content string) {
	result, err := GetResultByContent(content)
	if err != nil {
		general.CloseProgram(err)
	}
	fmt.Println(result)
}
