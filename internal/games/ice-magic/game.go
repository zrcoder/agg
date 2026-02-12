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
		pkg.WithChapters(g.chapters, g.reset),
		pkg.WithScene(sceneName, g.mainView),
	)
	g.Base = base
	g.reset()
	return g
}

func (g *Game) reset() {
	chapter, level := g.Base.ChapterIndex(), g.Base.LevelIndex()
	g.fires = 0
	g.failed = false
	g.parseGrid(chapter, level)
}

func (g *Game) done() bool {
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
