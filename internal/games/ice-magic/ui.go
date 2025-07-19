package icemagic

import "github.com/zrcoder/amisgo/comp"

func (g *Game) View() any {
	trs := make([]comp.Tr, len(g.grid))
	for i, line := range g.grid {
		tds := make([]comp.Td, len(line))
		for j := range line {
			tds[j] = line[j].View()
		}
		trs[i] = g.App.Tr().Tds(tds...).Height("30px")
	}
	return g.App.Wrapper().ClassName("w-1/2").Body(
		g.App.TableView().Padding(0).Trs(trs...),
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
