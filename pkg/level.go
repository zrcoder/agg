package pkg

import (
	"fmt"
	"strconv"
	"strings"

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
	Value string `json:"value"` // inner matained, dont't chage or use
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

func makeChapterLevelOptionValue(chapter, level int) string {
	return fmt.Sprintf("%d:%d", chapter, level)
}

func calChapterLevelIndex(optionValue string) (chapter, level int) {
	arr := strings.SplitN(optionValue, ":", 2)
	chapter, _ = strconv.Atoi(arr[0])
	level, _ = strconv.Atoi(arr[1])
	return
}
