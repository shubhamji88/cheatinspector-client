package fs

import (
	"fmt"
	"io"
	"os"
)

// CreateFile a new file with given name
func CreateFile(path string) bool {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return false
		}
		defer file.Close()
	}

	return true
}

// WriteFile contents to given file
func WriteFile(path string, data string) bool {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return false
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if isError(err) {
		return false
	}

	err = file.Sync()
	if isError(err) {
		return false
	}

	return true
}

// ReadFile and return the contents of the file
func ReadFile(path string) (bool, string) {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return false, ""
	}
	defer file.Close()

	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		// Break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// Break if error occured
		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}

	return true, string(text)
}

// DeleteFile the given file
func DeleteFile(path string) bool {
	// delete file
	var err = os.Remove(path)
	if isError(err) {
		return false
	}

	return true
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
