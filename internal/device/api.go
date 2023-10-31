package device

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/YashKumarVerma/cheatinspector-client/internal/config"
)

// GetDeviceDetailAPI fetches the device details from the remote API Server
func GetDeviceDetailAPI(deviceID string) (bool, Device) {
	resp, err := http.Get(config.Load.Server + "/device/" + deviceID)
	var device Device

	// check if non 200 response
	if err != nil {
		log.Fatalf("An Error Encountered %v", err)
	}
	defer resp.Body.Close()

	// read response body as required
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var parsedAPIResponse getDeviceDetailsAPIResponse
	json.Unmarshal([]byte(body), &parsedAPIResponse)

	return parsedAPIResponse.Err, device
}

// RegisterDeviceAPI registers a new device with the API server
func RegisterDeviceAPI(device registerDeviceStruct, teamID string) (bool, Device) {
	var deviceInfo Device

	// structure the data that needs to be sent
	postBody, _ := json.Marshal(map[string]string{
		"machineID":       device.MachineID,
		"teamId":          teamID,
		"friendlyName":    device.Name,
		"operatingSystem": device.OS,
		"frequency":       strconv.Itoa(device.Frequency),
	})
	responseBody := bytes.NewBuffer(postBody)

	// make the request
	resp, err := http.Post(config.Load.Server+"/device", "application/json", responseBody)

	// check if non 200 response
	if err != nil {
		log.Fatalf("An Error Encountered %v", err)
	}
	defer resp.Body.Close()

	// read response body as required
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return true, deviceInfo
	}

	var parsedAPIResponse getDeviceDetailsAPIResponse
	json.Unmarshal([]byte(body), &parsedAPIResponse)
	deviceInfo = parsedAPIResponse.Payload

	return parsedAPIResponse.Err, deviceInfo
}
