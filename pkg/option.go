package pkg

import (
	"strconv"
)

type Option func(*Base)

func WithScene(sceneName string, sceneFn func() any) Option {
	return func(b *Base) {
		b.sceneName = sceneName
		b.wsPath = "/ws/" + sceneName
		b.sceneFn = sceneFn
	}
}

func WithLevels(levels []Level, reset func()) Option {
	return func(b *Base) {
		b.levels = levels
		b.levelOptions = make([]any, len(levels))
		for i := range levels {
			b.levels[i].Value = strconv.Itoa(i)
			b.levelOptions[i] = &b.levels[i]
		}
		b.reset = reset
	}
}

func WithChapters(chapters []Chapter, reset func()) Option {
	return func(b *Base) {
		b.chapters = chapters
		for i := range b.chapters {
			for j := range b.chapters[i].Children {
				b.chapters[i].Children[j].Value = makeChapterLevelOptionValue(i, j)
			}
		}
		b.chapterOptions = make([]any, len(chapters))
		for i := range b.chapters {
			b.chapterOptions[i] = &b.chapters[i]
		}
		b.reset = reset
	}
}
