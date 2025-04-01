package hanoi

import (
	"fmt"

	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/schema"
)

const (
	sceneName = "hanoi"
)

func (g *Game) UI() comp.Page {
	return g.App.Page().
		Title(g.App.Tpl().Tpl("Tower of Hanoi").ClassName("text-2xl font-bold")).
		Body(
			g.App.Form().ColumnCount(3).
				Submit(func(s schema.Schema) error {
					return g.codeFn(s.Get("code").(string))
				}).
				Body(
					g.Game.Service(),
					g.App.Editor().Name("code"),
				),
		)
}

func (g *Game) Main() any {
	return g.Game.Main(
		g.IsDone(),
		fmt.Sprintf("Steps: %d, minimum: %d", g.steps, g.MinSteps()),
		"Press the pile key to select a pile.",
		g.pilesUI(),
	)
}

func (g *Game) pilesUI() comp.Flex {
	return g.App.Flex().Items(g.PileA.UI(), g.PileB.UI(), g.PileC.UI())
}

func (p *Pile) UI() comp.Form {
	done := p.Game.IsDone()
	var top any
	switch {
	case done && p.Index == len(p.Game.piles)-1:
		top = p.Game.starUI()
	case p.Game.ShiftDisk != nil && p.Game.ShiftDisk.Pile == p:
		top = p.Game.ShiftDisk.UI()
	default:
		top = p.Game.placeholderDiskUI()
	}
	n := len(p.Disks)
	disks := make([]any, n)
	for i := range disks {
		disks[i] = p.App.Flex().Items(p.Disks[i].UI())
	}
	return p.Game.pileForms[p.Index].Body(
		top,
		p.App.Service().Body(disks),
		p.App.Divider(),
		p.App.Button().ActionType("submit").HotKey(string(rune('A'+p.Index))).Disabled(done),
		p.App.Tpl().Tpl(string(rune('A'+p.Index))),
	).ClassName("w-36 h-auto mx-2 py-0 border-b-2")
}

func (g *Game) makePileForms() {
	g.pileForms = make([]comp.Form, PileCount)
	for i := 0; i < PileCount; i++ {
		i := i
		g.pileForms[i] = g.App.Form().Mode("inline").WrapWithPanel(false).Submit(
			func(s schema.Schema) error {
				g.SelectPile(g.piles[i])
				return nil
			})
	}
}

func (d *Disk) UI() comp.Shape {
	return d.Game.shape((float64(d.ID+1) * 20), "rectangle", d.Game.colors[d.ID])
}

const diskHeight = 40

func (g *Game) placeholderDiskUI() comp.Shape {
	return g.shape(diskHeight, "rectangle", "transparent")
}

func (g *Game) starUI() comp.Shape {
	return g.shape(diskHeight, "star", "orange")
}

func (g *Game) shape(width float64, shape, color string) comp.Shape {
	return g.App.Shape().ShapeType(shape).Width(width).Height(diskHeight).Color(color)
}
