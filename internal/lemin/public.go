package lemin

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// RunProgramWithFile - path is filepath, writes result to output. Close program if has error.
func RunProgramWithFile(path string, showContent bool) {
	err := WriteResultByFilePath(os.Stdout, path, showContent)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		os.Exit(1)
	}
}

// WriteResultByFilePath - path is filepath.
func WriteResultByFilePath(w io.Writer, path string, writeContent bool) error {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("WriteResultByFilePath: %w", err)
	} else if fInfo, _ := file.Stat(); fInfo.IsDir() {
		err = fmt.Errorf("%v is directory", fInfo.Name())
		return fmt.Errorf("WriteResultByFilePath: %w", err)
	}

	scanner := bufio.NewScanner(file)
	result, err := GetResult(scanner)
	if err != nil {
		return fmt.Errorf("WriteResultByFilePath: %w", err)
	}
	if writeContent {
		file.Seek(0, io.SeekStart)
		_, err = io.Copy(w, file)
		if err != nil {
			return fmt.Errorf("WriteResultByFilePath: %w", err)
		}
		fmt.Fprint(w, "\n\n# result\n")
	}
	result.WriteResult(w)

	return nil
}

// WriteResultByContent - using for Web, inputs writer for write result, writes nothing if returns error
func WriteResultByContent(w io.Writer, content string, writeContent bool) error {
	scanner := bufio.NewScanner(strings.NewReader(content))
	result, err := GetResult(scanner)
	if err != nil {
		return fmt.Errorf("WriteResultByContent: %w", err)
	}

	if writeContent {
		fmt.Fprintf(w, "%v\n\n", content)
	}
	result.WriteResult(w)

	return nil
}
