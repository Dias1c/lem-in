package lemin

import (
	"io/ioutil"
	"strings"
)

func getFileContent(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func splitToLines(content string) []string {
	if strings.Contains(content, "\r\n") {
		return strings.Split(content, "\r\n")
	}
	//Default
	return strings.Split(content, "\n")
}
