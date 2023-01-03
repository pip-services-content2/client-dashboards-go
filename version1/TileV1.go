package version1

type TileV1 struct {
	Title  string `json:"title"`
	Index  int    `json:"index"`
	Color  string `json:"color"`
	Size   string `json:"size"`
	Params any    `json:"params"`
}
