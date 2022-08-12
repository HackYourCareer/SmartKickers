package config

const (
	ServerAddress = "0.0.0.0:3000"

	TeamWhite = 1
	TeamBlue  = 2

	AttributeTeam   = "team"
	AttributeAction = "action"
)

var (
	WhiteTeamArea = [4]int{20, 21, 23, 25}
	BlueTeamArea  = [4]int{22, 24, 26, 27}
)
