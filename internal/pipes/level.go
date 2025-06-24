package pipes

type CellItem struct {
	Cell Cell
	X    int `json:"x"`
	Y    int `json:"y"`
}

type Level struct {
	Rows  int
	Cols  int
	Cells []CellItem
}
