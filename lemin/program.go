package lemin

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getFileContent(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func getTerrainFromLines(lines []string) (*terrain, error) {
	//
	startIdx := 0
	err := errors.New("getTerrainFromLines")
	countAnts := -1
	for i, line := range lines {
		startIdx = i + 1
		if strings.HasPrefix(line, "#") {
			continue
		} else {
			countAnts, err = strconv.Atoi(line)
			if err != nil {
				message := fmt.Sprintf("Incorrect value on count Ants: %q", line)
				return nil, errors.New(message)
			}
			break
		}

	}

	return nil, nil
}

func getCountAntsFromLine(lines []string) (int, []string, error) {
	err := errors.New("getCountAntsFromLine")
	startIdx := 0
	countAnts := -1
	for i, line := range lines {
		startIdx = i + 1
		if strings.HasPrefix(line, "#") {
			continue
		} else {
			countAnts, err = strconv.Atoi(line)
			if err != nil {
				message := fmt.Sprintf("Incorrect value on count Ants: %q", line)
				return 0, errors.New(message)
			}
			break
		}
	}
	return countAnts, lines[startIdx:], nil
}
