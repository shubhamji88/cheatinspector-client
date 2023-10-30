package team

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

// CreateOrJoinTeamScreen shows UI to select user choice on start-up
func CreateOrJoinTeamScreen() string {

	// declare choices for interface
	prompt := promptui.Select{
		Label: "Team Status",
		Items: []string{"Create a team", "Join a team"},
	}

	// extract data from selection
	_, result, err := prompt.Run()

	// if error in selection, show error and move to team creation
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "create"
	}

	// transform user response into strings which are later checked by other modules
	if result == "Create a team" {
		return "create"
	}

	// if not create, then join is the only second option.
	return "join"
}
