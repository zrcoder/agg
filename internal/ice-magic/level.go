package icemagic

import (
	"fmt"

	"github.com/zrcoder/agg/pkg"
)

const (
	levelsInEachChapter = 9
	totalLevels         = 10
)

const (
	sceneName = "ice-magic"

	Empty = ' '
	Wall  = '#'
	Fire  = 'f'
	Magic = 'm'
	Ice   = 'i'
)

type Level struct {
	HideMagicButton bool
	Grid            [][]byte
}

var imgdic = map[byte]string{
	Wall:  "/static/ice-magic/wall.svg",
	Fire:  "/static/ice-magic/fire.svg",
	Magic: "/static/ice-magic/magic.svg",
	Ice:   "/static/ice-magic/ice.svg",
}

func (g *Game) CurrentChapter() pkg.Chapter {
	return g.chapters[g.Base.ChapterIndex()]
}

func (g *Game) CurrentLevel() Level {
	chapter := g.chapters[g.Base.ChapterIndex()]
	return chapter.Children[g.Base.LevelIndex()].Data.(Level)
}

func (g *Game) initLevels() {
	g.chapters = []pkg.Chapter{
		{
			Children: []pkg.Level{
				{}, {}, {}, {}, {}, {}, {}, {}, {},
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
