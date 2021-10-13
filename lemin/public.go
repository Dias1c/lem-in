package lemin

import (
	"fmt"
	"io"
	"lem-in/general"
	"lem-in/lemin/anthill"
	"strings"
)

// RunProgramWithFile - path is filepath
func RunProgramWithFile(path string) {
	fileContent, err := getFileContent(path)
	if err != nil {
		general.CloseProgram(err)
	}
	err = PrintResultByContent(fileContent)
	if err != nil {
		general.CloseProgram(err)
	}
}

// Prints Result
func PrintResultByContent(content string) error {
	lines := splitToLines(content)
	myterrain, err := anthill.GetAnthillFromLines(lines)
	if err != nil {
		return errInvalidDataFormat(err)
	}
	err = myterrain.Validate()
	if err != nil {
		return err
	}
	err = myterrain.Match()
	if err != nil {
		return errPaths(err)
	}
	fmt.Println(content)
	myterrain.PrintResultByPaths()
	return nil
}

// GetResultByContent - Using for web.
// antsLimit - limit on the number of ants (0 - No restrictions)
func GetResultByContent(content string, antsLimit int) (string, error) {
	lines := splitToLines(content)
	myterrain, err := anthill.GetAnthillFromLines(lines)
	if err != nil {
		return "", errInvalidDataFormat(err)
	}
	if antsLimit != 0 && myterrain.AntsCount > antsLimit {
		return "", fmt.Errorf("limit on the number of ants %v", antsLimit)
	}
	err = myterrain.Validate()
	if err != nil {
		return "", err
	}
	err = myterrain.Match()
	if err != nil {
		return "", errPaths(err)
	}
	myterrain.GenerateResult()
	return fmt.Sprintf("%v\n\n%v", content, strings.Join(myterrain.Result, "\n")), nil
}

func WriteResultByContent(content string, w io.Writer) error {
	lines := splitToLines(content)
	myterrain, err := anthill.GetAnthillFromLines(lines)
	if err != nil {
		return errInvalidDataFormat(err)
	}
	err = myterrain.Validate()
	if err != nil {
		return err
	}
	err = myterrain.Match()
	if err != nil {
		return errPaths(err)
	}
	fmt.Fprintf(w, "%v\n\n", content)
	myterrain.WriteResult(w)
	return nil
}
