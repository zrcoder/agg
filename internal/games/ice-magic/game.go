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

func (g *Game) playerMoveLeft() error {
	leftCol, leftSprite := g.player.leftSprite()
	row := g.grid[g.player.Y]
	switch leftSprite.ID {
	case Blank:
		rowPre := row[:leftCol]
		rowSuf := row[leftCol+2:]
		if leftSprite.Width > 1 {
			leftSprite.Width--
			rowPre = append(rowPre, leftSprite, g.player)
		} else {
			rowPre = append(rowPre, g.player)
		}
		g.player.X--
		right := rowSuf[0]
		if right.ID == Blank {
			right.X--
			right.Width++
		} else {
			size1blank := &Sprite{ID: Blank, X: g.player.X + 1, Y: g.player.Y, Width: 1}
			rowPre = append(rowPre, size1blank)
		}
		rowPre = append(rowPre, rowSuf...)
		g.grid[leftSprite.Y] = rowPre
		return g.UpdateUI()
		// TODO: update ui
		g.checkUp(g.player.X+1, g.player.Y)
		g.checkDown(g.player.X, g.player.Y)
	case Ice:
	case Fire:
		g.failed = true
	case Wall:
	// TODO
	default:
	}
	return g.UpdateUI()
}

func (g *Game) playerMoveRight() error {
	rightCol, rightSprite := g.player.rightSprite()
	row := g.grid[g.player.Y]
	switch rightSprite.ID {
	case Blank:
		rowPre := row[:rightCol-1]
		rowSuf := row[rightCol:]
		if rightSprite.Width > 1 {
			rightSprite.Width--
			rowSuf = row[rightCol-1:] // insert player
		} else {
			row[rightCol] = g.player
		}
		g.player.X--
		left := row[rightCol-2]
		if left.ID == Blank {
			left.Width++
		} else {
			size1blank := &Sprite{ID: Blank, X: g.player.X - 1, Y: g.player.Y, Width: 1}
			rowPre = append(rowPre, size1blank)
		}
		rowPre = append(rowPre, rowSuf...)
		g.grid[rightSprite.Y] = rowPre
		return g.UpdateUI()
		// TODO: update ui
		g.checkUp(g.player.X+1, g.player.Y)
		g.checkDown(g.player.X, g.player.Y)
	case Ice:
	case Fire:
		g.failed = true
	case Wall:
	// TODO
	default:
	}
	return g.UpdateUI()
}

func (g *Game) checkUp(x, y int) {
	// TODO
}

func (g *Game) checkDown(x, y int) {
	// TODO
}
