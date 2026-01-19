package icemagic

import (
	"time"

	"github.com/zrcoder/amisgo/comp"
)

const (
	Blank  = ' '
	Wall   = '='
	Fire   = 'F'
	Player = 'M'
	Ice    = 'I'

	animateWaitMillisecond = 300
)

type Sprite struct {
	*Game
	ID         byte
	Left       *Sprite
	Right      *Sprite
	RowIndex   int
	X          int
	Y          int
	Width      int
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
	case Blank:
		td.Style(noBorderTdStyle).Body(".")
	}
	return td
}

func (s *Sprite) moveLeft() error {
	switch s.ID {
	case Player:
		return s.playerMoveLeft()
	case Ice:
		return s.iceSlideLeft()
	}
	return nil
}

func (s *Sprite) moveRight() error {
	switch s.ID {
	case Player:
		return s.playerMoveRight()
	case Ice:
		return s.iceSlideRight()
	}
	return nil
}

func (s *Sprite) playerMoveLeft() error {
	g := s.Game
	player := g.player
	left := player.Left
	var err error
	switch left.ID {
	case Blank:
		if err = left.hSwap(player); err != nil {
			return err
		}
		if err = left.checkUp(); err != nil {
			return err
		}
		return player.checkDown()
	case Ice:
		left.moveLeft()
	case Fire:
		g.failed = true
	case Wall:
	// TODO
	default:
	}
	return nil
}

func (s *Sprite) playerMoveRight() error {
	g := s.Game
	player := s.Game.player
	right := player.Right
	var err error
	switch right.ID {
	case Blank:
		if err = player.hSwap(right); err != nil {
			return err
		}
		if err = right.checkUp(); err != nil {
			return err
		}
		return player.checkDown()
	case Ice:
		right.moveRight()
	case Fire:
		g.failed = true
	case Wall:
	// TODO
	default:
	}
	return nil
}

func (s *Sprite) iceSlideLeft() error {
	left := s.Left
	if left == nil {
		return nil
	}
	switch left.ID {
	case Fire:
		s.Game.fires--
		left.ID = Blank
		s.ID = Blank
		return s.Game.UpdateUI()
	case Blank:
		err := left.hSwap(s)
		if err != nil {
			return err
		}
		// TODO, chack fallings
		return s.iceSlideLeft()
	}
	return nil
}
func (s *Sprite) iceSlideRight() error {
	return nil
}

func (s *Sprite) hSwap(o *Sprite) error {
	left, right := s.Left, o.Right
	row := s.Game.grid[s.Y]
	c1, c2 := s.RowIndex, o.RowIndex
	row[c1] = o
	row[c2] = s
	s.RowIndex, o.RowIndex = c2, c1
	s.X, o.X = o.X, s.X
	s.Left = o
	s.Right = right
	o.Right = s
	o.Left = left
	if left != nil {
		left.Right = o
	}
	if right != nil {
		right.Left = s
	}
	time.Sleep(animateWaitMillisecond * time.Millisecond)
	return s.Game.UpdateUI()
}

func (s *Sprite) checkDown() error {
	if s == nil {
		return nil
	}
	switch s.ID {
	case Player, Ice:
		down := s.downSprite()
		if down == nil {
			return nil
		}
		err := s.vSwap(down)
		if err != nil {
			return err
		}
		return s.checkDown()
	}
	return nil // TODO
}

func (s *Sprite) checkUp() error {

	return nil // TODO
}

func (s *Sprite) downSprite() *Sprite {
	if s.Y == len(s.Game.grid)-1 {
		return nil
	}
	nextRow := s.Game.grid[s.Y+1]
	for _, ns := range nextRow {
		if ns.X == s.X {
			return ns
		}
	}
	return nil
}

func (s *Sprite) vSwap(o *Sprite) error {
	// TODO
	return s.Game.UpdateUI()
}
