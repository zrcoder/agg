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
)

var (
	stepTime = 100 * time.Millisecond
	flagType = map[byte]string{
		Blank:  "blank",
		Wall:   "wall",
		Fire:   "fire",
		Player: "player",
		Ice:    "ice",
	}
)

type Sprite struct {
	*Game
	TypeFlag   byte
	Type       string // just for debug and log
	X          int
	Y          int
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
	td := s.Td().Colspan(1).Width("40px")
	switch s.TypeFlag {
	case Wall:
		td.Background("#E9967A").Style(s.borderStyle())
	case Fire:
		td.Body("🔥").Align("center").Style(noBorderTdStyle)
	case Player:
		td.Body("😺").Align("center").Style(noBorderTdStyle)
	case Ice:
		td.Background("#87CEFA").Style(s.borderStyle())
	case Blank:
		td.Style(noBorderTdStyle)
	}
	return td
}

func (s *Sprite) borderStyle() map[string]int {
	style := map[string]int{}
	if s.LeftFixed {
		style["borderLeftWidth"] = 0
	}
	if s.RightFixed {
		style["borderRightWidth"] = 0
	}
	return style
}

func (s *Sprite) moveLeft() bool {
	switch s.TypeFlag {
	case Player:
		return s.playerMoveLeft()
	case Ice:
		return s.iceSlideLeft()
	}
	return false
}

func (s *Sprite) moveRight() bool {
	switch s.TypeFlag {
	case Player:
		return s.playerMoveRight()
	case Ice:
		return s.iceSlideRight()
	}
	return false
}

func (s *Sprite) playerMoveLeft() bool {
	g := s.Game
	player := g.player
	left := player.Left()
	if left == nil {
		return false
	}
	switch left.TypeFlag {
	case Blank:
		if ok := g.hSwap(left, player); !ok {
			return false
		}
		g.checkUpsFall(left)
		return player.fall()
	case Ice:
		left.iceSlideLeft()
	case Fire:
		g.failed = true
	case Wall:
	// TODO
	default:
	}
	return false
}

func (s *Sprite) playerMoveRight() bool {
	g := s.Game
	player := s.Game.player
	right := player.Right()
	switch right.TypeFlag {
	case Blank:
		if ok := g.hSwap(player, right); !ok {
			return false
		}
		g.checkUpsFall(right.Up())
		return player.fall()
	case Ice:
		right.moveRight()
	case Fire:
		g.failed = true
	case Wall:
	// TODO
	default:
	}
	return false
}

func (s *Sprite) iceSlideLeft() bool {
	left := s.Left()
	if left == nil {
		return false
	}
	switch left.TypeFlag {
	case Fire:
		s.Game.fires--
		left.TypeFlag = Blank
		s.TypeFlag = Blank
		time.Sleep(stepTime)
		err := s.Game.UpdateUI()
		if err != nil {
			return false
		}
		s.Game.checkUpsFall(left.Up())
		s.Game.checkUpsFall(s.Up())
		return true
	case Blank:
		up := s.Up()
		defer func() {
			s.Game.checkUpsFall(up)
		}()
		if ok := s.Game.hSwap(left, s); !ok {
			return false
		}
		if ok := s.fall(); !ok {
			return s.iceSlideLeft()
		}
		return false
	}
	return false
}
func (s *Sprite) iceSlideRight() bool {
	right := s.Right()
	if right == nil {
		return false
	}
	switch right.TypeFlag {
	case Fire:
		up1 := s.Up()
		up2 := right.Up()
		s.Game.fires--
		right.TypeFlag = Blank
		s.TypeFlag = Blank
		time.Sleep(stepTime)
		err := s.Game.UpdateUI()
		if err != nil {
			return false
		}
		s.Game.checkUpsFall(up1)
		s.Game.checkUpsFall(up2)
		return true
	case Blank:
		up := s.Up()
		defer func() {
			s.Game.checkUpsFall(up)
		}()
		if ok := s.Game.hSwap(s, right); !ok {
			return false
		}
		if ok := s.fall(); !ok {
			return s.iceSlideRight()
		}
		return false
	}
	return false
}

func (s *Sprite) downSprite() *Sprite {
	if s.Y == len(s.Game.grid)-1 {
		return nil
	}
	nextRow := s.Game.grid[s.Y+1]
	return nextRow[s.X]
}

func (s *Sprite) Left() *Sprite {
	if s.X == 0 {
		return nil
	}
	return s.Game.grid[s.Y][s.X-1]
}
func (s *Sprite) Right() *Sprite {
	row := s.Game.grid[s.Y]
	n := len(row)
	if s.X == n-1 {
		return nil
	}
	return row[s.X+1]
}
func (s *Sprite) Up() *Sprite {
	if s.Y == 0 {
		return nil
	}
	return s.Game.grid[s.Y-1][s.X]
}
func (s *Sprite) Down() *Sprite {
	if s.Y == len(s.Game.grid)-1 {
		return nil
	}
	return s.Game.grid[s.Y+1][s.X]
}

func (s *Sprite) fall() bool {
	if s == nil {
		return false
	}
	g := s.Game
	res := false
	switch s.TypeFlag {
	case Player:
		for y := s.Y; y < len(g.grid)-1; y++ {
			down := s.downSprite()
			switch down.TypeFlag {
			case Blank:
				ok := g.vSwap(s, down)
				if !ok {
					return false
				}
				res = true
			case Fire:
				g.failed = true
				time.Sleep(stepTime)
				return g.UpdateUI() == nil
			}
		}
	case Ice:
		for y := s.Y; y < len(g.grid)-1; y++ {
			down := s.downSprite()
			switch down.TypeFlag {
			case Blank:
				ok := s.Game.vSwap(s, down)
				if !ok {
					return false
				}
				res = true
			case Fire:
				s.fires--
				s.TypeFlag = Blank
				down.TypeFlag = Blank
				return s.Game.UpdateUI() == nil
			}
		}
	case Fire:
		for y := s.Y; y < len(g.grid)-1; y++ {
			down := s.downSprite()
			switch down.TypeFlag {
			case Blank:
				ok := s.Game.vSwap(s, down)
				if !ok {
					return false
				}
			}
			res = true
		}
	}
	return res
}
