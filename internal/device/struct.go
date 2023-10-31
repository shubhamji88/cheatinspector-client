package device

//
// Device contains all details of given device
//
type Device struct {
	Frequency int    `json:"frequency"`
	Name      string `json:"friendlyName"`
	MachineID string `json:"machineID"`
	OS        string `json:"operatingSystem"`
	Timestamp int64  `json:"timestamp"`
}

//
// getDeviceDetailsAPIResponse
//
type getDeviceDetailsAPIResponse struct {
	Err     bool   `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Payload Device `json:"payload"`
}

//
// registerDeviceStruct
//
type registerDeviceStruct struct {
	Name      string `json:"friendlyName"`
	MachineID string `json:"machineID"`
	TeamId    string `json:"teamId"`
	OS        string `json:"operatingSystem"`
	Frequency int    `json:"frequency"`
}
