package fs

import "time"

// FileDetails contain all information about any given file
type FileDetails struct {
	Name         string
	Path         string
	Size         int64
	LineCount    int
	LastModified time.Time
	Contents     string
}
