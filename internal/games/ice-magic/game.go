package icemagic

import (
	"time"

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
	g.parseGrid(chapter, level)
}

func (g *Game) Done() bool {
	return g.fires == 0
}
func (g *Game) checkUpsFall(s *Sprite) {
	if s == nil {
		return
	}
	x := s.X
	for y := s.Y; y >= 0; y-- {
		s = g.grid[y][x]
		if !s.fall() {
			return
		}
	}
}

func (g *Game) hSwap(src, dst *Sprite) bool {
	if src.Y != dst.Y {
		return false
	}
	row := g.grid[src.Y]
	row[src.X], row[dst.X] = dst, src
	src.X, dst.X = dst.X, src.X
	time.Sleep(stepTime)
	err := g.UpdateUI()
	return err == nil
}

func (g *Game) vSwap(src, dst *Sprite) bool {
	if src.X != dst.X {
		return false
	}
	x := src.X
	sRow := g.grid[src.Y]
	oRow := g.grid[dst.Y]
	sRow[x], oRow[x] = dst, src
	src.Y, dst.Y = dst.Y, src.Y
	time.Sleep(stepTime)
	err := g.UpdateUI()
	return err == nil
}
