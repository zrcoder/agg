package hanoi

import (
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/schema"
)

const (
	sceneName = "hanoi"
)

func (g *Game) UI() comp.Page {
	return g.App.Page().
		Title(g.Tpl().Tpl("Tower of Hanoi").ClassName("text-2xl font-bold")).
		Body(
			g.Form().AutoFocus(true).ColumnCount(2).WrapWithPanel(false).
				Submit(func(s schema.Schema) error {
					return g.codeFn(s.Get("code").(string))
				}).
				Body(
					g.Flex().Items(g.topUI()),
					g.Wrapper().ClassName("w-1/2").Body(g.Game.Service()),
					g.Wrapper().ClassName("w-1/2").Body(g.App.Editor().Size("xxl").Language("go").Name("code")),
					g.Flex().Justify("center").Items(g.levelUI(), g.Wrapper(), g.Button().Label("Go").Icon("fa fa-play").ActionType("submit").HotKey("ctrl+g")),
				),
		)
}

func (g *Game) topUI() comp.Tpl {
	if g.IsDone() {
		return g.SucceedUI()
	}
	return g.StateUI(g.State())
}

func (g *Game) Main() any {
	return g.pilesUI()
}

func (g *Game) pilesUI() comp.Flex {
	return g.App.Flex().Items(g.PileA.UI(), g.Wrapper(), g.PileB.UI(), g.Wrapper(), g.PileC.UI())
}

func (g *Game) levelUI() comp.Flex {
	return g.App.Flex().Items(
		g.PrevForm,
		g.App.Tpl().Tpl(g.CurrentLevel().Name).ClassName("text-xl font-bold text-info pr-3"),
		g.NextForm,
	)
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
	m := maxDiskCount - n
	disks := make([]comp.Flex, maxDiskCount)
	for i := 0; i < m; i++ {
		disks[i] = p.App.Flex().Items(p.Game.placeholderDiskUI()).ClassName("h-10 py-0 my-0")
	}
	for i := 0; i < n; i++ {
		disks[i+m] = p.App.Flex().Items(p.Disks[i].UI()).ClassName("h-10 py-0 my-0")
	}
	return p.Game.pileForms[p.Index].Body(
		top,
		p.App.Service().Body(disks),
		p.App.Divider(),
		p.Flex().Items(p.App.Tpl().Tpl(string(rune('A'+p.Index)))),
	).ClassName("w-36 h-auto")
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

const diskHeight = 40

func (d *Disk) UI() comp.Shape {
	return d.Game.shape((float64(d.ID+1) * diskHeight), "rectangle", d.Game.colors[d.ID])
}

func (g *Game) placeholderDiskUI() comp.Shape {
	return g.shape(diskHeight, "rectangle", "transparent")
}

func (g *Game) starUI() comp.Shape {
	return g.shape(diskHeight, "star", "orange")
}

func (g *Game) shape(width float64, shape, color string) comp.Shape {
	return g.App.Shape().ShapeType(shape).Width(width).Height(diskHeight).Color(color).ClassName("py-0 my-0")
}
