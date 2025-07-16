package icemagic

import (
	"bytes"
	"fmt"

	"github.com/zrcoder/agg/internal/ice-magic/levels"
	"github.com/zrcoder/agg/pkg"
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

type Game struct {
	*amisgo.App
	*pkg.Base
	grid [][]byte
}

func New(app *amisgo.App) *Game {
	g := &Game{
		App: app,
	}
	levels := make([]pkg.Level, totalLevels)
	for i := range levels {
		chapter, section := calChapterSection(i)
		levels[i].Name = fmt.Sprintf("%d-%d", chapter, section)
		levels[i].Value = i
	}
	base := pkg.New(
		app,
		pkg.WithLevels(levels, g.Reset),
		pkg.WithScene(sceneName, g.Main),
	)
	g.Base = base
	g.Reset()
	return g
}

func (g *Game) Reset() {
	chapter, section := calChapterSection(g.CurrentLevel().Value)
	data, err := levels.FS.ReadFile(fmt.Sprintf("%d/%d.txt", chapter, section))
	if err != nil {
		panic(err)
	}
	g.grid = bytes.Split(data, []byte{'\n'})
}

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

func (g *Game) Done() bool {
	return false
}
