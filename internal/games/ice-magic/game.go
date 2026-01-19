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
	player   *Sprite
	failed   bool // failed if the play is burned
	fires    int
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
	g.fires = 0
	g.failed = false
	g.grid = g.parseGrid(chapter, level)
}

func (g *Game) Done() bool {
	return g.fires == 0
}

func (g *Game) checkUp(x, y int) {
	// TODO
}

func (g *Game) checkDown(x, y int) {
	// TODO
}
