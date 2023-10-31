package sensor

import "runtime"

// Init stars the sensor calls to collect data
func Init() {
	Load.MachineID = getMachineId()
	Load.OS = runtime.GOOS
	Load.Frequency = 10
}

// Load returns current sensor data for other modules
var Load Sensor
