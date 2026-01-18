package icemagic

import (
	"github.com/zrcoder/agg/internal/games/ice-magic/levels"
	"github.com/zrcoder/amisgo/comp"
)

const (
	Blank  = ' '
	Wall   = '='
	Fire   = 'F'
	Player = 'M'
	Ice    = 'I'
)

type Sprite struct {
	*Game
	X          int
	Y          int
	Width      int
	ID         byte
	LeftFixed  bool
	RightFixed bool
}

type Position struct {
	Row, Col int
}

func (s *Sprite) Fixed() bool {
	return s.LeftFixed || s.RightFixed
}

var noBorderTdStyle = map[string]int{
	"borderLeftWidth":   0,
	"borderRightWidth":  0,
	"borderTopWidth":    0,
	"borderBottomWidth": 0,
}

func (s *Sprite) View() comp.Td {
	td := s.Td().Colspan(s.Width).Width("40px")
	switch s.ID {
	case Wall:
		td.Background("#E9967A")
	case Fire:
		td.Body("🔥").Align("center").Style(noBorderTdStyle)
	case Player:
		td.Body("😺").Align("center").Style(noBorderTdStyle)
	case Ice:
		style := map[string]int{}
		if !s.LeftFixed {
			style["borderLeftWidth"] = 0
		}
		if !s.RightFixed {
			style["borderRightWidth"] = 0
		}
		td.Background("#87CEFA").Style(style)
	default:
		td.Style(noBorderTdStyle)
	}
	return td
}

func (s *Sprite) leftSprite() (int, *Sprite) {
	if s.X == 0 {
		return 0, nil
	}
	row := s.Game.grid[s.Y]
	i, left := 0, row[0]
	for left.X+left.Width < s.X {
		i++
		left = row[i]
	}
	return i, left
}

func (s *Sprite) rightSprite() (int, *Sprite) {
	if s.X+s.Width == levels.Cols {
		return 0, nil
	}
	row := s.Game.grid[s.Y]
	i := len(row) - 1
	right := row[i]
	for right.X > s.X+s.Width {
		i--
		right = row[i]
	}
	return i, right
}
