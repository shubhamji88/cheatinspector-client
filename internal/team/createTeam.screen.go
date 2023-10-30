package team

import (
	"errors"
	"fmt"
	"os/user"

	"github.com/manifoldco/promptui"
)

//  getTeamIdScreen to get the team ID input from user
func getTeamIdScreen() string {

	// ensure that team-id is atleast 4 characters long
	validate := func(input string) error {
		if len(input) < 4 {
			return errors.New("team id should be at-least 4 characters long")
		}
		return nil
	}

	//  assign the team name as the current user's username
	var teamID string
	u, err := user.Current()
	if err == nil {
		teamID = u.Username
	}

	// since team id is used to join a team, ensure that they are protected and hidden in UI
	prompt := promptui.Prompt{
		Label:    "Team ID",
		Validate: validate,
		Mask:     '*',
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return teamID
	}

	return result
}

//  getTeamNameScreen to get the team name input from user
func getTeamNameScreen() string {

	// ensure that team names are at-least 6 characters long.
	validate := func(input string) error {
		if len(input) < 6 {
			return errors.New("friendly team name should be at-least 6 characters")
		}
		return nil
	}

	// get current username
	var teamID string
	u, err := user.Current()
	if err == nil {
		teamID = u.Username
	}

	//  show the prompt and return data to user
	prompt := promptui.Prompt{
		Label:    "Team Name",
		Validate: validate,
		Default:  "team " + teamID,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return teamID
	}

	return result
}

// CreateTeamScreen shows UI to select user choice on start-up
func CreateTeamScreen() CreateTeamStruct {

	teamDetails := CreateTeamStruct{
		TeamID:   getTeamIdScreen(),
		TeamName: getTeamNameScreen(),
	}

	return teamDetails
}
