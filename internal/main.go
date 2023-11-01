package main

import (
	"fmt"
	"time"

	"github.com/YashKumarVerma/cheatinspector-client/internal/ably"
	"github.com/YashKumarVerma/cheatinspector-client/internal/config"

	"github.com/YashKumarVerma/cheatinspector-client/internal/device"
	"github.com/YashKumarVerma/cheatinspector-client/internal/fs"
	"github.com/YashKumarVerma/cheatinspector-client/internal/sensor"

	"github.com/YashKumarVerma/cheatinspector-client/internal/team"
	"github.com/YashKumarVerma/cheatinspector-client/internal/watchman"
)

func main() {
	// initialize all modules
	config.Init()
	sensor.Init()
	watchman.Init()

	// load all sensor data
	userDeviceInfo := sensor.Load
	ably.Init(userDeviceInfo.MachineID)

	// holders for user data
	var UserTeam team.Team
	var UserDevice device.Device

	notFound, deviceInfo := device.GetDeviceDetailAPI(userDeviceInfo.MachineID)
	teamOperation := team.CreateOrJoinTeamScreen()
	if teamOperation == "create" {
		teamDetails := team.CreateTeamScreen()
		teamAPIResponse := team.CreateTeamAPI(teamDetails)
		UserTeam = teamAPIResponse
	} else {
		teamDetails := team.JoinTeamScreen()
		teamAPIResponse := team.GetTeamDetailAPI(teamDetails.TeamID)
		UserTeam = teamAPIResponse
	}
	//  if not found, then initiate login process
	if notFound == true {
		//register or join a team


		fmt.Println("\nDevice not registered on cheatInspector, registration in process...")
		deviceDetails := device.CreateTeamScreen()
		teamNotFound, deviceAPIResponse := device.RegisterDeviceAPI(deviceDetails, UserTeam.ID)

		if teamNotFound == true {
			fmt.Println("Team does not exist")
		} else {
			fmt.Println("Device registered.")
			UserDevice = deviceAPIResponse
		}
	} else {
		fmt.Println("\nDevice already registered.")
		UserDevice = deviceInfo
	}

	success, folderNames := fs.ListFolders("./")
	if !success {
		fmt.Errorf("error reading folder names")
		return
	}

	fmt.Println("Device ID: ", sensor.Load.MachineID)

	// Updating presence
	ably.UserOnlinePresence(userDeviceInfo.MachineID)
	defer ably.UserLeavePresence()
	// repeat the process based on config.frequency
	for i := 1; i >= 0; i++ {
		for _, folder := range folderNames {
			filesNotIgnored, _ := watchman.IndexAllFiles(folder)
			for _, i := range filesNotIgnored {
				_, fileDetails := fs.AnalyzeFile(i)
				watchman.ProcessFile(fileDetails)
			}
		}
		time.Sleep(3 * time.Second)
		watchman.ResetAggregator()
		watchman.ResetTotal()
	}

	_ = UserDevice
}
