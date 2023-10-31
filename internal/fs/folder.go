package fs

import (
	"io/ioutil"
	"log"
	"os"
)

// ListAll returns list of all files or folders in given directory
func ListAll(path string) (bool, []string) {
	var fileList []string
	file, err := os.Open(path)
	if err != nil {
		return false, fileList
	}

	defer file.Close()

	names, err := file.Readdirnames(0)
	if err != nil {
		return false, fileList
	}
	fileList = names

	return true, fileList
}

// ListFolders returns list of folders only
func ListFolders(path string) (bool, []string) {
	var folderNames []string
	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
		return false, folderNames
	}

	for _, f := range files {
		if f.IsDir() {
			folderNames = append(folderNames, f.Name())
		}
	}

	return true, folderNames
}

// ListFiles returns list of files only
func ListFiles(path string) (bool, []string) {
	var fileNames []string
	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
		return false, fileNames
	}

	for _, f := range files {
		if !f.IsDir() {
			fileNames = append(fileNames, f.Name())
		}
	}

	return true, fileNames
}
