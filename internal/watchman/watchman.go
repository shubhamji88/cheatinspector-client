package watchman

import (
	"os"
	"path/filepath"

	"github.com/YashKumarVerma/cheatinspector-client/internal/fs"
	"github.com/sergi/go-diff/diffmatchpatch"
)

// Watchman is responsible for the following actions
// - walking the directory tree starting from the current working directory.
// - index all files along with their metadata and save into in-memory cache
// - sync data with cloud provider
// - save last snapshot of cache locally

// IndexAllFiles walks the passed path, and returns the list of ignored and
// focused files
func IndexAllFiles(path string) (filesNotIgnored []string, filesIgnored []string) {
	var focusedFiles []string
	var ignoredFiles []string

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if isKnownConfigs(info.Name()) {
			ignoredFiles = append(ignoredFiles, path)
		} else if childOfIgnoredDirectory(path) {
			ignoredFiles = append(ignoredFiles, path)
		} else if info.IsDir() {
			ignoredFiles = append(ignoredFiles, path)
		} else {
			focusedFiles = append(focusedFiles, path)
		}
		return nil
	})

	return focusedFiles, ignoredFiles
}

// ProcessFile calculates the diff between subsequent calls
func ProcessFile(file fs.FileDetails) bool {
	_, newDetails := fs.AnalyzeFile(file.Path)
	var difference Diff
	diffCalculator := diffmatchpatch.New()
	newDiffCalculator := diffmatchpatch.New()

	// check if entry exist in hashmap
	if oldDetails, ok := index[file.Path]; ok {
		diff := diffCalculator.DiffMain(oldDetails.Contents, newDetails.Contents, true)
		difference.size = oldDetails.Size - newDetails.Size
		difference.timestamp = newDetails.LastModified.Sub(oldDetails.LastModified)
		difference.changes = diffCalculator.DiffLevenshtein(diff)

	} else {
		newDiff := newDiffCalculator.DiffMain(newDetails.Contents, "", true)
		difference.size = newDetails.Size
		difference.timestamp = newDetails.LastModified.Sub(newDetails.LastModified)
		difference.changes = newDiffCalculator.DiffLevenshtein(newDiff)
	}

	// Aggregate the entropy to calculate net change per cycle
	Aggregate(difference.changes)
	Total(int64(len(newDetails.Contents)))

	// save the current state of file into cache
	setCache(newDetails)

	return true
}
