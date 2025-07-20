package pkg

import (
	"github.com/zrcoder/amisgo/comp"
)

type Chapter struct {
	Label    string  `json:"label"`
	Children []Level `json:"children"`
	Data     any     `json:"-"`
}

type Level struct {
	Label string `json:"label"`
	Data  any    `json:"-"`
	Value any    `json:"value"` // Value is a inner field
}

type LevelMeta struct {
	Chapter int `json:"chapter"`
	Level   int `json:"level"`
}

func (b *Base) LevelIndex() int {
	return b.levelIndex
}

func (b *Base) ChapterIndex() int {
	return b.chapterIndex
}

func (b *Base) LevelUI() comp.Flex {
	return b.App.Flex().Items(
		b.LeveSelectForm,
		b.ResetForm,
	)
}
