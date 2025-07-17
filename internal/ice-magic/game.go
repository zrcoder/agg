package icemagic

import (
	"bytes"
	"fmt"

	"github.com/zrcoder/agg/internal/ice-magic/levels"
	"github.com/zrcoder/agg/pkg"
	"github.com/zrcoder/amisgo"
)

type Game struct {
	*amisgo.App
	*pkg.Base
	chapters []pkg.Chapter
	grid     [][]byte
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
	data, err := levels.FS.ReadFile(fmt.Sprintf("%d/%d.txt", chapter+1, level+1))
	if err != nil {
		panic(err)
	}
	g.grid = bytes.Split(data, []byte{'\n'})
}

func (g *Game) Done() bool {
	return false
}
