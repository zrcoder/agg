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

func (g *Game) playerMoveLeft() bool {
	player := g.player
	left := player.Left()
	if left == nil {
		return false
	}
	switch left.Kind {
	case Blank:
		up := player.Up()
		if ok := g.swap(left, player, stepTime); !ok {
			return false
		}
		g.checkFall(up)
		return player.fall()
	case Ice:
		if !left.iceSlideLeft() {
			return g.player.climbLeft()
		}
		return false
	case Fire:
		player.PlayerDie()
		g.UpdateUI()
	case Wall:
		return g.player.climbLeft()
	default:
	}
	return false
}

func (g *Game) playerMoveRight() bool {
	player := g.player
	right := player.Right()
	switch right.Kind {
	case Blank:
		up := player.Up()
		if ok := g.swap(player, right, stepTime); !ok {
			return false
		}
		g.checkFall(up)
		return player.fall()
	case Ice:
		if !right.iceSlideRight() {
			return g.player.climbRight()
		}
		return false
	case Fire:
		player.PlayerDie()
		g.UpdateUI()
	case Wall:
		return g.player.climbRight()
	default:
	}
	return false
}

func (g *Game) checkFall(s *Sprite) {
	if s == nil || s.Kind == Blank || s.Kind == Wall {
		return
	}
	bar := s.Bar()
	g.FallBars(bar)
}

func (g *Game) swap(src, dst *Sprite, duration time.Duration) bool {
	if !g.swapQuietly(src, dst) {
		return false
	}
	time.Sleep(duration)
	return g.UpdateUI() == nil
}

func (g *Game) swapQuietly(src, dst *Sprite) bool {
	if src == nil || dst == nil {
		return false
	}
	sRow := g.grid[src.Y]
	dRow := g.grid[dst.Y]
	sRow[src.X], dRow[dst.X] = dst, src
	src.X, dst.X = dst.X, src.X
	src.Y, dst.Y = dst.Y, src.Y
	return true
}

func (g *Game) FallBars(bars ...*Bar) bool {
	if len(bars) == 0 {
		return false
	}
	var upBars, newBars []*Bar
	for _, b := range bars {
		if !b.CanFall() {
			continue
		}
		upBars = append(upBars, b.GetUpFallBars()...)
		newBars = append(newBars, b.FallBar1StepQuietly()...)
	}
	time.Sleep(stepTime)
	g.UpdateUI()
	g.FallBars(newBars...)
	g.FallBars(upBars...)
	return true
}
