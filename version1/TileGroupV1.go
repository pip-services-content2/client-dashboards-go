package version1

type TileGroupV1 struct {
	Title string    `json:"title"`
	Index int       `json:"index"`
	Tiles []*TileV1 `json:"tiles"`
}
