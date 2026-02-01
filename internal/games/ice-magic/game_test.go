package icemagic

import (
	"testing"

	"github.com/zrcoder/amisgo"
)

func TestNew(t *testing.T) {
	app := amisgo.New()

	game := New(app)

	if game == nil {
		t.Fatal("New() returned nil")
	}
	if game.App == nil {
		t.Error("New() did not set App")
	}
	if game.Base == nil {
		t.Error("New() did not set Base")
	}
	if len(game.chapters) == 0 {
		t.Error("New() did not initialize chapters")
	}
}

func TestReset(t *testing.T) {
	app := amisgo.New()
	game := New(app)

	game.Reset()

	if len(game.grid) == 0 {
		t.Error("Reset() did not create grid")
	}
	if game.player == nil {
		t.Error("Reset() did not set player")
	}
	if game.fires == 0 {
		t.Error("Reset() did not set fires")
	}
	if game.failed {
		t.Error("Reset() set failed to true")
	}
}

func TestDone(t *testing.T) {
	app := amisgo.New()
	game := New(app)
	game.Reset()

	if game.Done() {
		t.Error("Done() returned true when fires > 0")
	}

	game.fires = 0
	if !game.Done() {
		t.Error("Done() returned false when fires == 0")
	}
}

func TestHSwapQuite(t *testing.T) {
	app := amisgo.New()
	game := &Game{App: app}

	game.grid = [][]*Sprite{
		{{TypeFlag: Blank, X: 0, Y: 0, Game: game}, {TypeFlag: Blank, X: 1, Y: 0, Game: game}},
		{{TypeFlag: Blank, X: 0, Y: 1, Game: game}, {TypeFlag: Blank, X: 1, Y: 1, Game: game}},
	}

	src := game.grid[0][0]
	dst := game.grid[0][1]

	result := game.hSwapQuite(src, dst)

	if !result {
		t.Error("hSwapQuite() returned false")
	}
	if game.grid[0][0] != dst || game.grid[0][1] != src {
		t.Error("hSwapQuite() did not swap sprites")
	}
	if src.X != 1 || dst.X != 0 {
		t.Error("hSwapQuite() did not update X coordinates")
	}
}

func TestHSwapQuiteDifferentRows(t *testing.T) {
	app := amisgo.New()
	game := &Game{App: app}

	game.grid = [][]*Sprite{
		{{TypeFlag: Blank, X: 0, Y: 0, Game: game}, {TypeFlag: Blank, X: 1, Y: 0, Game: game}},
		{{TypeFlag: Blank, X: 0, Y: 1, Game: game}, {TypeFlag: Blank, X: 1, Y: 1, Game: game}},
	}

	src := game.grid[0][0]
	dst := game.grid[1][0]

	result := game.hSwapQuite(src, dst)

	if result {
		t.Error("hSwapQuite() returned true for different rows")
	}
}

func TestVSwapQuite(t *testing.T) {
	app := amisgo.New()
	game := &Game{App: app}

	game.grid = [][]*Sprite{
		{{TypeFlag: Blank, X: 0, Y: 0, Game: game}, {TypeFlag: Blank, X: 1, Y: 0, Game: game}},
		{{TypeFlag: Blank, X: 0, Y: 1, Game: game}, {TypeFlag: Blank, X: 1, Y: 1, Game: game}},
	}

	src := game.grid[0][0]
	dst := game.grid[1][0]

	result := game.vSwapQuite(src, dst)

	if !result {
		t.Error("vSwapQuite() returned false")
	}
	if game.grid[0][0] != dst || game.grid[1][0] != src {
		t.Error("vSwapQuite() did not swap sprites")
	}
	if src.Y != 1 || dst.Y != 0 {
		t.Error("vSwapQuite() did not update Y coordinates")
	}
}

func TestVSwapQuiteDifferentColumns(t *testing.T) {
	app := amisgo.New()
	game := &Game{App: app}

	game.grid = [][]*Sprite{
		{{TypeFlag: Blank, X: 0, Y: 0, Game: game}, {TypeFlag: Blank, X: 1, Y: 0, Game: game}},
		{{TypeFlag: Blank, X: 0, Y: 1, Game: game}, {TypeFlag: Blank, X: 1, Y: 1, Game: game}},
	}

	src := game.grid[0][0]
	dst := game.grid[0][1]

	result := game.vSwapQuite(src, dst)

	if result {
		t.Error("vSwapQuite() returned true for different columns")
	}
}

func TestSpriteLeft(t *testing.T) {
	app := amisgo.New()
	game := &Game{App: app}

	game.grid = [][]*Sprite{
		{{TypeFlag: Blank, X: 0, Y: 0, Game: game}, {TypeFlag: Blank, X: 1, Y: 0, Game: game}},
	}

	sprite := game.grid[0][1]
	left := sprite.Left()

	if left == nil {
		t.Error("Left() returned nil")
	}
	if left != game.grid[0][0] {
		t.Error("Left() returned wrong sprite")
	}

	sprite = game.grid[0][0]
	left = sprite.Left()
	if left != nil {
		t.Error("Left() from edge returned non-nil")
	}
}

func TestSpriteRight(t *testing.T) {
	app := amisgo.New()
	game := &Game{App: app}

	game.grid = [][]*Sprite{
		{{TypeFlag: Blank, X: 0, Y: 0, Game: game}, {TypeFlag: Blank, X: 1, Y: 0, Game: game}},
	}

	sprite := game.grid[0][0]
	right := sprite.Right()

	if right == nil {
		t.Error("Right() returned nil")
	}
	if right != game.grid[0][1] {
		t.Error("Right() returned wrong sprite")
	}

	sprite = game.grid[0][1]
	right = sprite.Right()
	if right != nil {
		t.Error("Right() from edge returned non-nil")
	}
}

func TestSpriteUp(t *testing.T) {
	app := amisgo.New()
	game := &Game{App: app}

	game.grid = [][]*Sprite{
		{{TypeFlag: Blank, X: 0, Y: 0, Game: game}},
		{{TypeFlag: Blank, X: 0, Y: 1, Game: game}},
	}

	sprite := game.grid[1][0]
	up := sprite.Up()

	if up == nil {
		t.Error("Up() returned nil")
	}
	if up != game.grid[0][0] {
		t.Error("Up() returned wrong sprite")
	}

	sprite = game.grid[0][0]
	up = sprite.Up()
	if up != nil {
		t.Error("Up() from top returned non-nil")
	}
}

func TestSpriteDown(t *testing.T) {
	app := amisgo.New()
	game := &Game{App: app}

	game.grid = [][]*Sprite{
		{{TypeFlag: Blank, X: 0, Y: 0, Game: game}},
		{{TypeFlag: Blank, X: 0, Y: 1, Game: game}},
	}

	sprite := game.grid[0][0]
	down := sprite.Down()

	if down == nil {
		t.Error("Down() returned nil")
	}
	if down != game.grid[1][0] {
		t.Error("Down() returned wrong sprite")
	}

	sprite = game.grid[1][0]
	down = sprite.Down()
	if down != nil {
		t.Error("Down() from bottom returned non-nil")
	}
}

func TestSpriteLeftUp(t *testing.T) {
	app := amisgo.New()
	game := &Game{App: app}

	game.grid = [][]*Sprite{
		{{TypeFlag: Blank, X: 0, Y: 0, Game: game}, {TypeFlag: Blank, X: 1, Y: 0, Game: game}},
		{{TypeFlag: Blank, X: 0, Y: 1, Game: game}, {TypeFlag: Blank, X: 1, Y: 1, Game: game}},
	}

	sprite := game.grid[1][1]
	leftUp := sprite.LeftUp()

	if leftUp == nil {
		t.Error("LeftUp() returned nil")
	}
	if leftUp != game.grid[0][0] {
		t.Error("LeftUp() returned wrong sprite")
	}

	sprite = game.grid[0][0]
	leftUp = sprite.LeftUp()
	if leftUp != nil {
		t.Error("LeftUp() from corner returned non-nil")
	}
}

func TestSpriteRightUp(t *testing.T) {
	app := amisgo.New()
	game := &Game{App: app}

	game.grid = [][]*Sprite{
		{{TypeFlag: Blank, X: 0, Y: 0, Game: game}, {TypeFlag: Blank, X: 1, Y: 0, Game: game}},
		{{TypeFlag: Blank, X: 0, Y: 1, Game: game}, {TypeFlag: Blank, X: 1, Y: 1, Game: game}},
	}

	sprite := game.grid[1][0]
	rightUp := sprite.RightUp()

	if rightUp == nil {
		t.Error("RightUp() returned nil")
	}
	if rightUp != game.grid[0][1] {
		t.Error("RightUp() returned wrong sprite")
	}

	sprite = game.grid[0][1]
	rightUp = sprite.RightUp()
	if rightUp != nil {
		t.Error("RightUp() from corner returned non-nil")
	}
}
