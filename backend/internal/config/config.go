package config

const (
	ServerAddress = "0.0.0.0:3000"

	// Aliases for team ID's received from node application
	TeamWhite = 1
	TeamBlue  = 2

	// Constants for query attributes from API URLs
	AttributeTeam   = "team"
	AttributeAction = "action"

	// Constant for heatmap accuracy, the larger the more accurate but more resource demanding
	HeatmapAccuracy = 100
)

// Numbers corresponding to areas on the table, as received from node application
var (
	WhiteTeamArea = [4]int{20, 21, 23, 25}
	BlueTeamArea  = [4]int{22, 24, 26, 27}
)
