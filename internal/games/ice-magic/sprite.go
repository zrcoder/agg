package icemagic

import "github.com/zrcoder/amisgo/comp"

const (
	Blank  = ' '
	Wall   = '#'
	Fire   = 'f'
	Player = 'm'
	Ice    = 'i'
)

type Sprite struct {
	*Game
	Width           int
	ID              byte
	BackgroundColor string
	LeftFixed       bool
	RightFixed      bool
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
