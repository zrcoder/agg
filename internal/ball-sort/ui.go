package ballsort

import (
	"fmt"

	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/schema"
)

const (
	sceneName = "ball-sort"
	ballSize  = 40
	Title     = "Ball Sort Puzzle"
)

func (g *Game) Main() any {
	return g.Game.Main(
		g.IsDone(),
		fmt.Sprintf("Done: %d/%d", g.DoneBottlesCount, g.CurrentLevel().Value),
		"Click any bottle or press the bottle key to select a bottle.",
		g.bottlesUI(),
	)
}

func (g *Game) bottlesUI() any {
	bottles := make([]any, len(g.Bottles))
	for i, bottle := range g.Bottles {
		bottles[i] = bottle.UI()
	}
	return bottles
}

func (b *Bottle) UI() any {
	done := b.IsDone()
	var top any
	switch {
	case done:
		top = b.Game.starUI()
	case b.IsShiftBall():
		top = b.Game.ShiftBall.UI()
	default:
		top = b.Game.placeholderBallUI()
	}

	items := make([]any, BottleBallCount)
	for i, ball := range b.Balls {
		items[BottleBallCount-i-1] = ball.UI()
	}
	for i := BottleBallCount - len(b.Balls) - 1; i >= 0; i-- {
		items[i] = b.Game.placeholderBallUI()
	}

	key := string(rune('A' + b.Index))
	return b.Game.bottleForms[b.Index].Body(
		b.App.Wrapper().ClassName("mx-2").Body(top),
		b.App.Wrapper().ClassName("relative w-18 h-auto mx-2").Body(
			items,
			b.Button().HotKey(key).ActionType("submit").Reload(sceneName).
				ClassName("absolute inset-0 h-full rounded-lg bottle-button").Disabled(done),
		),
		b.Flex().Items(b.Tpl().Tpl(key)),
	)
}

func (b *Ball) UI() comp.Shape {
	return b.Game.shape("circle", b.Game.colors[b.Type])
}

func (g *Game) placeholderBallUI() comp.Shape {
	return g.shape("circle", "transparent")
}

func (g *Game) starUI() comp.Shape {
	return g.shape("star", "orange")
}

func (g *Game) shape(shape, color string) comp.Shape {
	return g.App.Shape().ShapeType(shape).Width(ballSize).Height(ballSize).Color(color)
}

func (g *Game) makeBottleForms() {
	g.bottleForms = make([]comp.Form, len(g.colors)+EmptyBottles)
	for i := range g.bottleForms {
		i := i
		g.bottleForms[i] = g.Form().WrapWithPanel(false).Submit(
			func(s schema.Schema) error {
				g.SelectBottle(i)
				return g.UpdateUI()
			},
		)
	}
}
