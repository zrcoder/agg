package icemagic

import "github.com/zrcoder/amisgo/comp"

func (g *Game) View() any {
	var trs = make([]comp.Tr, len(g.grid))
	for i, line := range g.grid {
		var tds = make([]comp.Td, len(line))
		for j := range line {
			var view any
			if img, ok := imgdic[line[j]]; ok {
				view = g.App.Image().Src(img).ImageMode("original").InnerClassName("border-none")
			}
			tds[j] = g.App.Td().Align("center").Body(view).Width("20px").Padding(0)
		}
		trs[i] = g.App.Tr().Height("10px").Tds(tds...)
	}
	return g.App.Wrapper().ClassName("w-1/2").Body(
		g.App.TableView().Padding(0).Border(false).Trs(trs...),
	)
}

func (g *Game) Main() any {
	return g.Base.Main(
		g.Done(),
		"",
		"",
		g.View(),
	)
}
