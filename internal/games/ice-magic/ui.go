package icemagic

import (
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/schema"
)

func (g *Game) view() any {
	if g.failed {
		return g.App.Tpl().Tpl("Failed!").ClassName("text-2xl font-bold text-red-500 ")
	}
	return g.App.Wrapper().ClassName("w-2/3").Body(
		g.App.TableView().Border(false).Trs(
			g.App.Tr().Tds(
				g.App.Td().Body(g.sceneView()),
				g.App.Td().Body(g.buttonsPanel()),
			),
		),
	)
}

func (g *Game) sceneView() comp.TableView {
	trs := make([]comp.Tr, len(g.grid))
	for i, line := range g.grid {
		tds := make([]comp.Td, len(line))
		for j := range line {
			tds[j] = line[j].View()
		}
		trs[i] = g.App.Tr().Tds(tds...).Height("38px")
	}
	return g.App.TableView().BorderColor("white").Padding(0).Trs(trs...)
}

func (g *Game) buttonsPanel() comp.Wrapper {
	return g.App.Wrapper().Body(
		g.App.Flex().Items(
			g.buttonForm("J", "←", func() error {
				g.player.moveLeft()
				return nil
			}),
			g.App.Wrapper(),
			g.buttonForm("L", "→", func() error {
				g.player.moveRight()
				return nil
			}),
		),
		g.App.Wrapper(),
		g.App.Flex().Items(
			g.buttonForm("A", "↙", func() error {
				g.player.magicLeft()
				return nil
			}),
			g.App.Wrapper(),
			g.buttonForm("D", "↘", func() error {
				g.player.magicRight()
				return nil
			}),
		),
	)
}

func (g *Game) buttonForm(key, label string, action func() error) comp.Form {
	return g.App.Form().WrapWithPanel(false).Body(
		g.App.SubmitAction().Level("dark").HotKey(key).Label(label),
		g.App.Flex().Items(
			g.App.Tpl().Text(key),
		),
	).Submit(func(s schema.Schema) error {
		return action()
	})
}

func (g *Game) mainView() any {
	return g.Base.Main(
		g.done(),
		"",
		"",
		g.view(),
	)
}
