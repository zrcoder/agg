package ballsort

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
	if len(game.colors) != 8 {
		t.Errorf("New() did not set colors, got %v", len(game.colors))
	}
}

func TestReset(t *testing.T) {
	app := amisgo.New()
	game := New(app)

	game.Reset()

	if len(game.Bottles) == 0 {
		t.Error("Reset() did not create bottles")
	}
	if game.ShiftBall != nil {
		t.Error("Reset() did not reset ShiftBall")
	}
	if game.DoneBottlesCount != 0 {
		t.Errorf("Reset() did not reset DoneBottlesCount, got %v", game.DoneBottlesCount)
	}
}

func TestBottlePush(t *testing.T) {
	app := amisgo.New()
	game := &Game{App: app}
	bottle := &Bottle{Game: game, Balls: make([]*Ball, 0, BottleBallCount)}

	ball := &Ball{Type: 1, Bottle: bottle}
	bottle.Push(ball)

	if len(bottle.Balls) != 1 {
		t.Errorf("Push() did not add ball, got %v", len(bottle.Balls))
	}
	if bottle.Balls[0] != ball {
		t.Error("Push() did not add the correct ball")
	}
	if ball.Bottle != bottle {
		t.Error("Push() did not update ball's Bottle")
	}
}

func TestBottlePop(t *testing.T) {
	app := amisgo.New()
	game := &Game{App: app}
	bottle := &Bottle{Game: game, Balls: make([]*Ball, 0, BottleBallCount)}

	ball1 := &Ball{Type: 1, Bottle: bottle}
	ball2 := &Ball{Type: 2, Bottle: bottle}
	bottle.Push(ball1)
	bottle.Push(ball2)

	popped := bottle.Pop()

	if popped != ball2 {
		t.Errorf("Pop() returned wrong ball, got %v", popped)
	}
	if len(bottle.Balls) != 1 {
		t.Errorf("Pop() did not remove ball, got %v", len(bottle.Balls))
	}
	if bottle.Balls[0] != ball1 {
		t.Error("Pop() removed the wrong ball")
	}

	popped = bottle.Pop()
	if popped != ball1 {
		t.Errorf("Pop() returned wrong ball, got %v", popped)
	}
	if len(bottle.Balls) != 0 {
		t.Errorf("Pop() did not remove ball, got %v", len(bottle.Balls))
	}

	popped = bottle.Pop()
	if popped != nil {
		t.Errorf("Pop() from empty bottle returned %v, want nil", popped)
	}
}

func TestBottleTop(t *testing.T) {
	app := amisgo.New()
	game := &Game{App: app}
	bottle := &Bottle{Game: game, Balls: make([]*Ball, 0, BottleBallCount)}

	top := bottle.Top()
	if top != nil {
		t.Errorf("Top() from empty bottle returned %v, want nil", top)
	}

	ball1 := &Ball{Type: 1, Bottle: bottle}
	ball2 := &Ball{Type: 2, Bottle: bottle}
	bottle.Push(ball1)
	bottle.Push(ball2)

	top = bottle.Top()
	if top != ball2 {
		t.Errorf("Top() returned wrong ball, got %v", top)
	}
}

func TestBottleIsEmpty(t *testing.T) {
	app := amisgo.New()
	game := &Game{App: app}
	bottle := &Bottle{Game: game, Balls: make([]*Ball, 0, BottleBallCount)}

	if !bottle.IsEmpty() {
		t.Error("IsEmpty() returned false for empty bottle")
	}

	bottle.Push(&Ball{Type: 1, Bottle: bottle})
	if bottle.IsEmpty() {
		t.Error("IsEmpty() returned true for non-empty bottle")
	}
}

func TestBottleIsFull(t *testing.T) {
	app := amisgo.New()
	game := &Game{App: app}
	bottle := &Bottle{Game: game, Balls: make([]*Ball, 0, BottleBallCount)}

	if bottle.IsFull() {
		t.Error("IsFull() returned true for empty bottle")
	}

	for range BottleBallCount - 1 {
		bottle.Push(&Ball{Type: 1, Bottle: bottle})
	}
	if bottle.IsFull() {
		t.Error("IsFull() returned true for bottle with less than max balls")
	}

	bottle.Push(&Ball{Type: 1, Bottle: bottle})
	if !bottle.IsFull() {
		t.Error("IsFull() returned false for full bottle")
	}
}

func TestBottleCheckDone(t *testing.T) {
	app := amisgo.New()
	game := &Game{App: app}
	bottle := &Bottle{Game: game, Balls: make([]*Ball, 0, BottleBallCount)}

	result := bottle.checkDone()
	if result != 0 {
		t.Errorf("checkDone() returned %v for empty bottle, want 0", result)
	}

	for range BottleBallCount - 1 {
		bottle.Push(&Ball{Type: 1, Bottle: bottle})
	}
	result = bottle.checkDone()
	if result != 0 {
		t.Errorf("checkDone() returned %v for not full bottle, want 0", result)
	}

	bottle.Push(&Ball{Type: 2, Bottle: bottle})
	result = bottle.checkDone()
	if result != 0 {
		t.Errorf("checkDone() returned %v for mixed balls, want 0", result)
	}

	bottle.Balls = []*Ball{{Type: 1}, {Type: 1}, {Type: 1}, {Type: 1}}
	result = bottle.checkDone()
	if result != 1 {
		t.Errorf("checkDone() returned %v for same type balls, want 1", result)
	}
}

func TestBottleIsDone(t *testing.T) {
	app := amisgo.New()
	game := &Game{App: app}
	bottle := &Bottle{Game: game, Balls: make([]*Ball, 0, BottleBallCount)}

	if bottle.IsDone() {
		t.Error("IsDone() returned true for empty bottle")
	}

	for range BottleBallCount - 1 {
		bottle.Push(&Ball{Type: 1, Bottle: bottle})
	}
	if bottle.IsDone() {
		t.Error("IsDone() returned true for not full bottle")
	}

	bottle.Push(&Ball{Type: 2, Bottle: bottle})
	if bottle.IsDone() {
		t.Error("IsDone() returned true for mixed balls")
	}

	bottle.Balls = []*Ball{{Type: 1}, {Type: 1}, {Type: 1}, {Type: 1}}
	if !bottle.IsDone() {
		t.Error("IsDone() returned false for same type balls")
	}
}

func TestBottleIsShiftBall(t *testing.T) {
	app := amisgo.New()
	game := New(app)
	game.Reset()

	ball := game.Bottles[0].Top()
	if ball == nil {
		t.Fatal("Bottle has no balls")
	}

	game.ShiftBall = ball

	if !game.Bottles[0].IsShiftBall() {
		t.Error("IsShiftBall() returned false for bottle with shift ball")
	}

	if game.Bottles[1].IsShiftBall() {
		t.Error("IsShiftBall() returned true for bottle without shift ball")
	}
}

func TestSelectBottle(t *testing.T) {
	app := amisgo.New()
	game := New(app)
	game.Reset()

	if len(game.Bottles) == 0 {
		t.Fatal("Reset() did not create bottles")
	}

	game.SelectBottle(0)

	if game.ShiftBall == nil {
		t.Error("SelectBottle() did not set ShiftBall")
	}

	game.SelectBottle(1)

	if game.ShiftBall != nil {
		t.Error("SelectBottle() did not move ball to another bottle")
	}
}

func TestIsDone(t *testing.T) {
	app := amisgo.New()
	game := New(app)
	game.Reset()

	if game.IsDone() {
		t.Error("IsDone() returned true at start")
	}

	bottle := &Bottle{Game: game, Balls: make([]*Ball, 0, BottleBallCount)}
	for range BottleBallCount {
		bottle.Push(&Ball{Type: 1, Bottle: bottle})
	}
	game.Bottles[0] = bottle
	game.DoneBottlesCount = 1

	if game.IsDone() {
		t.Error("IsDone() returned true when not all bottles are done")
	}
}
