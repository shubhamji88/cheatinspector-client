package device

import (
	"errors"
	"fmt"
	"os/user"

	"github.com/YashKumarVerma/cheatinspector-client/internal/sensor"
	"github.com/manifoldco/promptui"
)

// to get friendly name for current device
func getDeviceNameScreen() string {
	validate := func(input string) error {
		if len(input) < 6 {
			return errors.New("friendly name should be at-least 6 characters long")
		}
		return nil
	}

	var teamID string
	u, err := user.Current()
	if err == nil {
		teamID = u.Username
	}

	prompt := promptui.Prompt{
		Label:    "Device Name",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return teamID
	}

	return result
}

// CreateOrJoinTeam shows UI to select user choice on start-up
func CreateTeamScreen() registerDeviceStruct {
	deviceDetails := registerDeviceStruct{
		Name:      getDeviceNameScreen(),
		MachineID: sensor.Load.MachineID,
		TeamId:    "",
		OS:        sensor.Load.OS,
		Frequency: sensor.Load.Frequency,
	}
	return deviceDetails
}
