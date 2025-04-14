package pkg

type Option func(*Base)

func WithLevels(levels []Level, reset func()) Option {
	return func(b *Base) {
		b.levels = levels
		b.reset = reset
	}
}

func WithScene(sceneName string, sceneFn func() any) Option {
	return func(b *Base) {
		b.sceneName = sceneName
		b.wsPath = "/ws/" + sceneName
		b.sceneFn = sceneFn
	}
}
