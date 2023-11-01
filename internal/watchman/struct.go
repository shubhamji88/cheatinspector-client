package watchman

import (
	"time"
)

//
// Diff contains the difference between cached data
//
type Diff struct {
	size      int64
	timestamp time.Duration
	changes   int
}

//
// entropyUpdateAPIStruct defines the data to be shared with server during polling the json tags have been purposely made smaller to minimize the data i/o
//
type entropyUpdateAPIStruct struct {
	device    string `json:"id"`
	entropy   int64  `json:"v"`
}
