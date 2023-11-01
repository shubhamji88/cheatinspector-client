package watchman

// aggregatorValue stores entropy of project
var TotalCalculator uint64

// Total is called by other routines to find total size of project
func Total(drop int64) {
	TotalCalculator += uint64(drop)
}

// ResetTotal sets the project size to zero, to be run after each cycle
func ResetTotal() {
	go UpdateSnapshot(TotalCalculator)

	TotalCalculator = 0
}
