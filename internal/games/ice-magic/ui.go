package icemagic

import (
	"fmt"

	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/schema"
)

func (g *Game) View() any {
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
		trs[i] = g.App.Tr().Tds(tds...).Height("30px")
	}
	return g.App.TableView().Padding(0).Trs(trs...)
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
				return nil
			}),
			g.App.Wrapper(),
			g.buttonForm("D", "↘", func() error {
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
		fmt.Println("Button triggered:", key)
		return action()
	})
}

func (g *Game) Main() any {
	return g.Base.Main(
		g.Done(),
		"",
		"",
		g.View(),
	)
}
