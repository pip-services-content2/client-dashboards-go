package version1

type DashboardV1 struct {
	// Identification
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	App    string `json:"app"`
	Kind   string `json:"kind"`

	// Content
	Groups []*TileGroupV1 `json:"groups"`
}
