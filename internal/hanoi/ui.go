package hanoi

import (
	"errors"

	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/schema"
)

const (
	sceneName = "hanoi"
)

func (g *Game) UI() comp.Page {
	return g.App.Page().
		Title(g.Tpl().Tpl("Tower of Hanoi").ClassName("text-2xl font-bold")).
		Toolbar(g.ThemeButtonGroupSelect()).
		Body(
			g.App.Service().Name(sceneName).InitFetch(true).
				GetData(func() (any, error) {
					return map[string]any{"main": g.Main()}, nil
				}).
				Ws(wsPath).Body(g.Amis().Name("main")),
		)
}

func (g *Game) Main() any {
	return g.Form().WrapWithPanel(false).ColumnCount(2).AutoFocus(true).
		Submit(
			func(s schema.Schema) error {
				code := s.Get("code")
				if code == nil {
					return errors.New("code is required")
				}
				go g.codeFn(code.(string))
				return nil
			},
		).Body(
		g.Flex().Items(g.stateUI()),
		g.Wrapper().ClassName("w-1/2").Body(g.pilesUI()),
		g.Wrapper().ClassName("w-1/2").Body(
			g.App.Editor().Size("xxl").Language("go").Name("code").
				Options(schema.Schema{
					"fontSize":         16,
					"wordWrap":         "on",
					"quickSuggestions": false,
				}).Value(g.CurrentLevel().Code),
		),
		g.Flex().ClassName("pt-4").Items(
			g.Button().Label("Go").Icon("fa fa-play").Primary(true).ActionType("submit").HotKey("ctrl+g"),
			g.Wrapper(),
			g.ResetForm,
			g.PrevForm,
			g.App.Tpl().Tpl(g.CurrentLevel().Name).ClassName("text-xl font-bold text-info pr-3"),
			g.NextForm,
		),
	)
}

func (g *Game) stateUI() comp.Tpl {
	if g.IsDone() {
		return g.SucceedUI()
	}
	return g.StateUI(g.State())
}

func (g *Game) pilesUI() comp.Service {
	return g.App.Service().Body(
		g.App.Flex().Items(g.App.Wrapper().ClassName("w-1/2").Body(g.PileC.UI())),
		g.App.Flex().Items(g.PileA.UI(), g.Wrapper(), g.PileB.UI()),
	)
}

func (p *Pile) UI() comp.TableView {
	trs := make([]comp.Tr, 0, maxDiskCount+1)

	var top comp.Tr
	switch {
	case p.Game.ShiftDisk != nil && p.Game.ShiftDisk.Pile == p:
		top = p.Game.ShiftDisk.UI()
	default:
		top = p.Game.placeholderDiskUI(true)
	}
	trs = append(trs, top)

	for m := maxDiskCount - len(p.Disks); m > 0; m-- {
		trs = append(trs, p.Game.placeholderDiskUI(false))
	}
	for i := len(p.Disks) - 1; i >= 0; i-- {
		trs = append(trs, p.Disks[i].UI())
	}

	key := string(rune('A' + p.Index))
	return p.TableView().Trs(
		p.Tr().Tds(
			p.Td().Style(p.tdBorderBottom()).Body(
				p.TableView().Trs(trs...),
			),
		),
		p.Tr().Tds(
			p.Td().Align("center").Style(p.tdBorderNone()).Body(
				p.App.Tpl().ClassName("text-xl font-bold").Tpl(key),
			),
		),
	)
}

func (d *Disk) UI() comp.Tr {
	tds := make([]comp.Td, 0, 2*maxDiskCount)
	blanks := maxDiskCount - d.ID - 1
	for i := 0; i < blanks; i++ {
		tds = append(tds, d.Td().Style(d.tdBorderNone()))
	}
	for i := 0; i < 2*(d.ID+1); i++ {
		tds = append(tds, d.Td().Background(d.colors[d.ID]).Style(d.tdBorderNone()))
	}
	for i := 0; i < blanks; i++ {
		tds = append(tds, d.Td().Style(d.tdBorderNone()))
	}
	return d.Tr().Tds(tds...)
}

func (g *Game) placeholderDiskUI(isTop bool) comp.Tr {
	tds := make([]comp.Td, maxDiskCount*2)
	for i := range tds {
		tds[i] = g.Td().Style(g.tdBorderNone())
		if !isTop && i == len(tds)/2 {
			tds[i].Style(g.tdBorderLeft())
		}
	}
	return g.Tr().Tds(tds...)
}

func (g *Game) tdBorderLeft() schema.Schema {
	return schema.Schema{
		"borderLeftWidth":   1,
		"borderRightWidth":  0,
		"borderTopWidth":    0,
		"borderBottomWidth": 0,
	}
}

func (g *Game) tdBorderBottom() schema.Schema {
	return schema.Schema{
		"borderLeftWidth":   0,
		"borderRightWidth":  0,
		"borderTopWidth":    0,
		"borderBottomWidth": 1,
	}
}

func (g *Game) tdBorderNone() schema.Schema {
	return schema.Schema{"borderWidth": 0}
}
