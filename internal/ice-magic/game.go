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

const (
	sceneName = "ice-magic"

	Empty = ' '
	Wall  = '#'
	Fire  = 'f'
	Magic = 'g'
	Ice   = 'i'
)

func New(app *amisgo.App) *Game {
	g := &Game{
		App: app,
	}
	levels := make([]pkg.Level, 2)
	for i := range levels {
		levels[i].Name = fmt.Sprintf("Level %d", i+1)
		levels[i].Value = i + 1
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
	level := g.CurrentLevel().Value
	chapter, err := levels.FS.ReadFile(fmt.Sprintf("%d.txt", level))
	if err != nil {
		panic(err)
	}
	g.grid = bytes.Split(chapter, []byte{'\n'})
}

var imgdic = map[byte]string{
	Wall:  "/static/ice-magic/wall.svg",
	Fire:  "/static/ice-magic/fire.svg",
	Magic: "/static/ice-magic/magic.svg",
	Ice:   "/static/ice-magic/ice.svg",
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
	return g.App.Wrapper().ClassName("w-7/12").Body(
		g.App.TableView().Padding(0).Border(false).Trs(trs...),
	)
}

func (g *Game) Main() any {
	return g.Base.Main(
		false,
		"Ice",
		"",
		g.View(),
	)
}
