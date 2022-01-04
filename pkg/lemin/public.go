package lemin

import (
	"io"

	lemin "github.com/Dias1c/lem-in/internal/lemin"
)

// RunProgramWithFile - path is filepath, writes result to output. Close program if has error.
func RunProgramWithFile(path string, showContent bool) {
	lemin.RunProgramWithFile(path, showContent)
}

// WriteResultByFilePath - path is filepath.
func WriteResultByFilePath(w io.Writer, path string, writeContent bool) error {
	return lemin.WriteResultByFilePath(w, path, writeContent)
}

// WriteResultByContent - using for Web, inputs writer for write result, writes nothing if returns error
func WriteResultByContent(w io.Writer, content string, writeContent bool) error {
	return lemin.WriteResultByContent(w, content, writeContent)
}
