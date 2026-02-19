package ballsort

import (
	"github.com/zrcoder/agg/pkg"
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

const (
	BottleBallCount = 4
	EmptyBottles    = 2
)

type Game struct {
	*amisgo.App
	*pkg.Base
	Bottles          []*Bottle
	ShiftBall        *Ball
	DoneBottlesCount int
	colors           []string
	bottleForms      []comp.Form
}

type Bottle struct {
	*Game
	Index int
	Balls []*Ball
}

type Ball struct {
	*Bottle
	Type int
}

func New(app *amisgo.App) *Game {
	g := &Game{
		App: app,
		colors: []string{
			"red", "#7FFF00", "blue", "yellow", "violet", "pink", "green", "orange",
		},
	}
	base := pkg.New(
		app,
		pkg.WithLevels(levels, g.Reset),
		pkg.WithScene(sceneName, g.Main),
	)
	g.Base = base
	g.makeBottleForms()
	g.Reset()
	return g
}

func (g *Game) Reset() {
	g.Base.Shuffle(len(g.colors), func(i, j int) {
		g.colors[i], g.colors[j] = g.colors[j], g.colors[i]
	})
	n := g.currentBallTubes()
	balls := make([]*Ball, 0, n*BottleBallCount)
	for i := range n {
		for range BottleBallCount {
			balls = append(balls, &Ball{Type: i})
		}
	}
	g.Base.Shuffle(len(balls), func(i, j int) {
		balls[i], balls[j] = balls[j], balls[i]
	})
	g.Bottles = make([]*Bottle, 0, n+EmptyBottles)
	i := 0
	for range n {
		bottle := &Bottle{Game: g}
		bottle.Balls = balls[i : i+BottleBallCount]
		for _, ball := range bottle.Balls {
			ball.Bottle = bottle
		}
		i += BottleBallCount
		g.Bottles = append(g.Bottles, bottle)
	}
	for range EmptyBottles {
		bottle := &Bottle{
			Game:  g,
			Balls: make([]*Ball, 0, BottleBallCount),
		}
		g.Bottles = append(g.Bottles, bottle)
	}
	g.Base.Shuffle(len(g.Bottles), func(i, j int) {
		g.Bottles[i], g.Bottles[j] = g.Bottles[j], g.Bottles[i]
	})
	for i, bottle := range g.Bottles {
		bottle.Index = i
	}
	g.ShiftBall = nil
	g.DoneBottlesCount = 0
	for _, bottle := range g.Bottles {
		if bottle.IsDone() {
			g.DoneBottlesCount++
		}
	}
}

func (g *Game) SelectBottle(i int) {
	if i < 0 || i >= len(g.Bottles) {
		return
	}
	bottle := g.Bottles[i]
	if bottle.IsDone() {
		return
	}
	if bottle.IsEmpty() {
		if g.ShiftBall == nil {
			return
		}
		bottle.Push(g.ShiftBall)
		g.ShiftBall = nil
		return
	}
	if g.ShiftBall == nil {
		g.ShiftBall = bottle.Pop()
		return
	}
	if g.ShiftBall.Bottle.Index != i &&
		(g.ShiftBall.Type != bottle.Top().Type || bottle.IsFull()) {
		g.ShiftBall.Bottle.Push(g.ShiftBall)
		g.ShiftBall = bottle.Pop()
		return
	}
	bottle.Push(g.ShiftBall)
	g.ShiftBall = nil
	g.DoneBottlesCount += bottle.checkDone()
}

func (g *Game) IsDone() bool {
	return g.DoneBottlesCount == g.currentBallTubes()
}

func (b *Bottle) Pop() *Ball {
	n := len(b.Balls)
	if n == 0 {
		return nil
	}
	res := b.Balls[n-1]
	b.Balls = b.Balls[:n-1]
	return res
}

func (b *Bottle) Push(ball *Ball) {
	b.Balls = append(b.Balls, ball)
	ball.Bottle = b
}

func (b *Bottle) Top() *Ball {
	if len(b.Balls) == 0 {
		return nil
	}
	return b.Balls[len(b.Balls)-1]
}

func (b *Bottle) IsEmpty() bool {
	return len(b.Balls) == 0
}

func (b *Bottle) IsFull() bool {
	return len(b.Balls) == BottleBallCount
}

func (b *Bottle) checkDone() int {
	if !b.IsFull() {
		return 0
	}
	for i := 1; i < len(b.Balls); i++ {
		if b.Balls[i].Type != b.Balls[0].Type {
			return 0
		}
	}
	return 1
}

func (b *Bottle) IsDone() bool {
	return b.checkDone() == 1
}

func (b *Bottle) IsShiftBall() bool {
	return b.Game.ShiftBall != nil && b.Game.ShiftBall.Bottle == b
}
