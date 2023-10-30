package team

// JoinTeamScreen shows UI
func JoinTeamScreen() JoinTeamStruct {

	teamDetails := JoinTeamStruct{
		TeamID: getTeamIdScreen(),
	}

	return teamDetails
}
