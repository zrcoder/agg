package icemagic

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

func calChapterSection(level int) (chapter, section int) {
	chapter = level/levelsInEachChapter + 1
	section = level%levelsInEachChapter + 1
	return
}

var imgdic = map[byte]string{
	Wall:  "/static/ice-magic/wall.svg",
	Fire:  "/static/ice-magic/fire.svg",
	Magic: "/static/ice-magic/magic.svg",
	Ice:   "/static/ice-magic/ice.svg",
}
