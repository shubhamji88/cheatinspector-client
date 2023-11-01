package watchman

// aggregatorValue stores entropy of project
var AggregatorValue uint64

// Aggregate is called by other routines to update the aggregatorValue
func Aggregate(entropy int) {
	AggregatorValue += uint64(entropy)
}

// ResetAggregator sets the project entropy to zero, to be run after each cycle
func ResetAggregator(){
	go NotifyBackend(AggregatorValue)
	AggregatorValue = 0
}
