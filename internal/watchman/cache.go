package watchman

import (
	"github.com/YashKumarVerma/cheatinspector-client/internal/fs"
)

var index map[string]fs.FileDetails

// Init to initialize internal key-value storage for directory details
func Init() {
	index = make(map[string]fs.FileDetails)
}

// setCache saves the key and vale into storage
func setCache(details fs.FileDetails) {
	index[details.Path] = details
}

// loadCache returns details about given file from last capture
func loadCache(path string) (fileDetails fs.FileDetails) {
	return index[path]
}
