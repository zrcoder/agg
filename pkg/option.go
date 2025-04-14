package pkg

type Option func(*Game)

func WithLevels(levels []Level, reset func()) Option {
	return func(g *Game) {
		g.levels = levels
		g.reset = reset
	}
}

func WithScene(sceneName string, sceneFn func() any) Option {
	return func(g *Game) {
		g.sceneName = sceneName
		g.wsPath = "/ws/" + sceneName
		g.sceneFn = sceneFn
	}
}
