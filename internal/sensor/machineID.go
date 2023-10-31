package sensor

import (
	"log"

	"github.com/denisbrodbeck/machineid"
)

// getMachineId returns a unique signature of each device
func getMachineId() string {
	id, err := machineid.ID()
	if err != nil {
		log.Fatal(err)
	}
	return id
}
