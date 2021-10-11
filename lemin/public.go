package lemin

import (
	"fmt"
	"lem-in/general"
	"lem-in/lemin/anthill"
)

// RunProgramWithFile - path is filepath
func RunProgramWithFile(path string) {
	fileContent, err := getFileContent(path)
	if err != nil {
		general.CloseProgram(err)
	}
	RunProgramWithContent(fileContent)
}

// RunProgramWithContent - content is filecontent
func RunProgramWithContent(content string) {
	result, err := GetResultByContent(content)
	if err != nil {
		general.CloseProgram(err)
	}
	result = "" // Remove this
	fmt.Print(result)
}

// GetResultByContent - Using for web.
func GetResultByContent(content string) (string, error) {
	lines := splitToLines(content)
	myterrain, err := anthill.GetAnthillFromLines(lines)
	if err != nil {
		return "", errInvalidDataFormat(err)
	}
	err = myterrain.Validate()
	if err != nil {
		return "", err
	}
	result, err := myterrain.Match()
	if err != nil {
		return "Incorrect", errPaths(err)
	}
	return fmt.Sprintf("%v\n\n#Result:\n%v", content, result), nil
}
