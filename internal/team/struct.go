package team

import "github.com/YashKumarVerma/cheatinspector-client/internal/device"

// CreateTeamStruct to encapsulate all team data
type CreateTeamStruct struct {
	TeamID   string
	TeamName string
}

// JoinTeamStruct contains the data required to join a team
type JoinTeamStruct struct {
	TeamID string
}

// Team contains all details of given team
type Team struct {
	Name    string          `json:"friendlyName"`
	ID      string          `json:"id"`
	Devices []device.Device `json:"devices"`
}

// createTeamAPIResponse represents the response from server when creating a team
type createTeamAPIResponse struct {
	Err     bool   `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Payload Team   `json:"payload"`
}

// getTeamDetailsAPIResponse represents the response from server when fetching team details
type getTeamDetailsAPIResponse struct {
	Err     bool   `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Payload Team   `json:"payload"`
}
