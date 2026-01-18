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

func (g *Game) parseGrid(chapter, level int) [][]*Sprite {
	data, err := levels.FS.ReadFile(fmt.Sprintf("%d/%d.txt", chapter+1, level+1))
	if err != nil {
		panic(err)
	}
	rows := bytes.Split(data, []byte{'\n'})
	grid := make([][]*Sprite, len(rows))
	for i, row := range rows {
		j := 0
		for j < len(row) {
			id := row[j]
			sprite := &Sprite{ID: id, X: j, Y: i, Width: 1, Game: g}
			switch id {
			case Fire:
				g.fires++
				j++
			case Player:
				g.player = sprite
				j++
			default:
				j++
				for ; j < len(row) && row[j] == id; j++ {
					sprite.Width++
				}
			}
			grid[i] = append(grid[i], sprite)
		}
	}
	return grid
}
