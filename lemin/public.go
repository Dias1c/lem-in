package lemin

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Dias1c/lem-in/general"
)

// RunProgramWithFile - path is filepath
func RunProgramWithFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		general.CloseProgram(err)
	} else if fInfo, _ := file.Stat(); fInfo.IsDir() {
		err = fmt.Errorf("%v is directory", fInfo.Name())
		general.CloseProgram(err)
	}
	scanner := bufio.NewScanner(file)
	result, err := getResult(scanner)
	if err != nil {
		file.Close()
		general.CloseProgram(err)
	}
	file.Seek(0, io.SeekStart)
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		file.Close()
		general.CloseProgram(err)
	}
	file.Close()
	fmt.Println()
	result.WriteResult(os.Stdout)
}

// WriteResultByContent - using for Web, inputs writer for write result, writes nothing if returns error
func WriteResultByContent(w io.Writer, content string) error {
	scanner := bufio.NewScanner(strings.NewReader(content))
	result, err := getResult(scanner)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "%v\n", content)
	result.WriteResult(w)
	return nil
}
