package ballsort

import "github.com/zrcoder/agg/pkg"

const (
	LevelEasy   = "EASY"
	LevelMedium = "MEDIUM"
	LevelHard   = "HARD"
	LevelExpert = "EXPERT"
)

var levels = []pkg.Level{
	{Label: LevelEasy, Data: 5},
	{Label: LevelMedium, Data: 6},
	{Label: LevelHard, Data: 7},
	{Label: LevelExpert, Data: 8},
}

func (g *Game) CurrentLevel() pkg.Level {
	return levels[g.Base.LevelIndex()]
}

func (g *Game) currentBallTubes() int {
	return levels[g.Base.LevelIndex()].Data.(int)
}
