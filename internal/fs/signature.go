package fs

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// AnalyzeFile file for all details
func AnalyzeFile(path string) (success bool, file FileDetails) {
	var fileDetails FileDetails
	meta, err := os.Stat(path)
	if err != nil {
		return false, fileDetails
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading contents")
	}

	// load metadata about file
	fileDetails.Path = path
	fileDetails.Name = meta.Name()
	fileDetails.Size = meta.Size()
	fileDetails.LastModified = meta.ModTime()
	fileDetails.LineCount = lineCounter(path)
	fileDetails.Contents = string(data)

	return true, fileDetails
}

// lineCounter returns number of lines in the given file
func lineCounter(path string) int {
	file, err := os.OpenFile(path, os.O_RDWR, 644)
	if err != nil {
		return -1
	}

	defer file.Close()
	buf := make([]byte, 32*1024)
	counter := 0
	lineSep := []byte{'\n'}

	for {
		c, err := file.Read(buf)
		counter += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return counter

		case err != nil:
			return counter
		}
	}
	return counter
}
