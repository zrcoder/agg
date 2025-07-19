package icemagic

import (
	"github.com/zrcoder/agg/pkg"
	"github.com/zrcoder/amisgo"
)

type Game struct {
	*amisgo.App
	*pkg.Base
	chapters []pkg.Chapter
	grid     [][]*Sprite
}

func New(app *amisgo.App) *Game {
	g := &Game{
		App: app,
	}
	g.initLevels()

	base := pkg.New(
		app,
		pkg.WithChapters(g.chapters, g.Reset),
		pkg.WithScene(sceneName, g.Main),
	)
	g.Base = base
	g.Reset()
	return g
}

func (g *Game) Reset() {
	chapter, level := g.Base.ChapterIndex(), g.Base.LevelIndex()
	g.grid = g.parseGrid(chapter, level)
}

func (g *Game) Done() bool {
	return false
}
