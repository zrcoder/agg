package pipes

type CellKind int

const (
	KindStart CellKind = iota
	KindEnd
	KindStraight
	KindCross
	KindX
	KindL
)

var rotateStates = []int{1, 1, 2, 1, 2, 4}

type RotateState int

type LightedState int

const (
	LightedNone LightedState = iota
	LightedAll
	LightedFirst
	LightedSecond
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Cell struct {
	Kind         CellKind     `json:"kind"`
	RottedState  RotateState  `json:"rotateState"`
	LightedState LightedState `json:"lightedState"`
}

type Game struct {
	Board [][]*Cell
}

func New(level Level) *Game {
	board := make([][]*Cell, level.Rows)
	for i := range board {
		board[i] = make([]*Cell, level.Cols)
	}
	for _, item := range level.Cells {
		board[item.Y][item.X] = &item.Cell
	}
	return &Game{Board: board}
}
