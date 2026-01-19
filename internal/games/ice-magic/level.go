package icemagic

import (
	"bytes"
	"fmt"

	"github.com/zrcoder/agg/internal/games/ice-magic/levels"
	"github.com/zrcoder/agg/pkg"
)

const (
	sceneName = "ice-magic"
)

type Level struct {
	HideMagicButton bool
}

func (g *Game) CurrentChapter() pkg.Chapter {
	return g.chapters[g.Base.ChapterIndex()]
}

func (g *Game) CurrentLevel() *Level {
	chapter := g.chapters[g.Base.ChapterIndex()]
	return chapter.Children[g.Base.LevelIndex()].Data.(*Level)
}

func (g *Game) initLevels() {
	g.chapters = []pkg.Chapter{
		{
			Children: []pkg.Level{
				{
					Data: &Level{HideMagicButton: true},
				}, {}, {}, {}, {}, {}, {}, {}, {},
			},
		},
		{
			Children: []pkg.Level{
				{},
			},
		},
	}
	for i := range g.chapters {
		g.chapters[i].Label = fmt.Sprintf("Chapter %d", i+1)
		for j := range g.chapters[i].Children {
			g.chapters[i].Children[j].Label = fmt.Sprintf("%d-%d", i+1, j+1)
		}
	}
}

func (g *Game) parseGrid(chapter, level int) {
	data, err := levels.FS.ReadFile(fmt.Sprintf("%d/%d.txt", chapter+1, level+1))
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(data, []byte{'\n'})
	g.grid = make([][]*Sprite, len(lines))
	for y, line := range lines {
		g.grid[y] = make([]*Sprite, len(line))
		for x := range line {
			id := line[x]
			sprite := &Sprite{TypeFlag: id, Type: flagType[id], X: x, Y: y, Game: g}
			switch id {
			case Blank:
			case Fire:
				g.fires++
			case Player:
				g.player = sprite
			case Ice:
				sprite.checkToFixLeft()
			case Wall:
				sprite.checkToFixLeft()
			}
			g.grid[y][x] = sprite
		}
	}
}

func (s *Sprite) checkToFixLeft() {
	left := s.Left()
	if s.X > 0 && (left.TypeFlag == Wall || left.TypeFlag == Ice) {
		s.LeftFixed = true
		left.RightFixed = true
	}
}
