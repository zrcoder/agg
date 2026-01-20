package icemagic

import (
	"time"

	"github.com/zrcoder/amisgo/comp"
)

const (
	Blank    = ' '
	Wall     = '='
	Fire     = 'F'
	Player   = 'M'
	Ice      = 'I'
	IceFixed = 'i'
)

var (
	stepTime = 100 * time.Millisecond
)

type Sprite struct {
	*Game
	TypeFlag   byte
	X          int
	Y          int
	LeftFixed  bool
	RightFixed bool
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
		td.Body(s.Tpl().Tpl("🔥").ClassName("text-2xl")).Align("center").Style(noBorderTdStyle)
	case Player:
		td.Body(s.Tpl().Tpl("☃︎").ClassName("text-2xl")).Align("center").Style(noBorderTdStyle)
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

func (s *Sprite) magicLeft() {
	if s.TypeFlag != Player {
		return
	}
	s.magic(s.LeftDown())

}
func (s *Sprite) magicRight() {
	if s.TypeFlag != Player {
		return
	}
	s.magic((s.RightDown()))
}
func (s *Sprite) magic(dst *Sprite) {
	if dst == nil {
		return
	}
	switch dst.TypeFlag {
	case Blank:
		dst.TypeFlag = Ice
		left := dst.Left()
		if left != nil && (left.TypeFlag == Ice || left.TypeFlag == Wall) {
			left.RightFixed = true
			dst.LeftFixed = true
		}
		right := dst.Right()
		if right != nil && (right.TypeFlag == Ice || right.TypeFlag == Wall) {
			right.LeftFixed = true
			dst.RightFixed = true
		}
		s.Game.UpdateUI()
	case Ice:
		dst.TypeFlag = Blank
		dst.LeftFixed = false
		dst.RightFixed = false
		left := dst.Left()
		if left != nil {
			left.RightFixed = false
		}
		right := dst.Right()
		if right != nil {
			right.LeftFixed = false
		}
		s.Game.UpdateUI()
		s.Game.checkUpsFall(s.Left())
	}
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
		up := player.Up()
		if ok := g.hSwap(left, player); !ok {
			return false
		}
		g.checkUpsFall(up)
		return player.fall()
	case Ice:
		if !left.iceSlideLeft() {
			return g.player.climbLeft()
		}
		return false
	case Fire:
		g.failed = true
		g.UpdateUI()
	case Wall:
		return g.player.climbLeft()
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
		up := player.Up()
		if ok := g.hSwap(player, right); !ok {
			return false
		}
		g.checkUpsFall(up)
		return player.fall()
	case Ice:
		if !right.iceSlideRight() {
			return g.player.climbRight()
		}
		return false
	case Fire:
		g.failed = true
		g.UpdateUI()
	case Wall:
		return g.player.climbRight()
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
			s.iceSlideLeft()
		}
		return true
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
		s.Game.fires--
		right.TypeFlag = Blank
		s.TypeFlag = Blank
		time.Sleep(stepTime)
		err := s.Game.UpdateUI()
		if err != nil {
			return false
		}
		s.Game.checkUpsFall(right.Up())
		s.Game.checkUpsFall(s.Up())
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
			s.iceSlideRight()
		}
		return true
	}
	return false
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
	if s == nil || s.TypeFlag == Blank {
		return false
	}
	g := s.Game
	switch s.TypeFlag {
	case Player:
		for y := s.Y; y < len(g.grid)-1; y++ {
			down := s.Down()
			switch down.TypeFlag {
			case Blank:
				ok := g.vSwap(s, down)
				if !ok {
					return false
				}
			case Fire:
				s.TypeFlag = Blank
				down.TypeFlag = Blank
				g.failed = true
				time.Sleep(stepTime)
				return g.UpdateUI() == nil
			}
		}
		return true
	case Ice:
		res := false
		for y := s.Y; y < len(g.grid)-1; y++ {
			x1, x2, ok := s.checkIceDown()
			if ok {
				res = true
				s.iceRowDown(x1, x2)
			}
		}
		return res
	case Fire:
		for y := s.Y; y < len(g.grid)-1; y++ {
			down := s.Down()
			switch down.TypeFlag {
			case Blank:
				ok := s.Game.vSwap(s, down)
				if !ok {
					return false
				}
			}
		}
		return true
	}
	return false
}

func (s *Sprite) checkIceDown() (int, int, bool) {
	if s.TypeFlag != Ice {
		return 0, 0, false
	}
	x1, x2 := s.X, s.X
	row := s.grid[s.Y]
	for x1 > 0 && row[x1].LeftFixed && row[x1-1].TypeFlag == Ice {
		x1--
	}
	if row[x1].LeftFixed { // ice left is fixed by wall or other sprites
		return 0, 0, false
	}
	for x2 < len(row)-1 && row[x2].RightFixed && row[x2+1].TypeFlag == Ice {
		x2++
	}
	if row[x2].RightFixed { // ice right is fixed by wall or other sprites
		return 0, 0, false
	}
	for x := x1; x <= x2; x++ {
		down := s.grid[s.Y+1][x]
		if down.TypeFlag != Blank && down.TypeFlag != Fire {
			return 0, 0, false
		}
	}
	return x1, x2, true
}

func (s *Sprite) iceRowDown(x1, x2 int) bool {
	row := s.grid[s.Y]
	nextRow := s.grid[s.Y+1]
	for x := x1; x <= x2; x++ {
		ice := row[x]
		down := nextRow[x]
		switch down.TypeFlag {
		case Fire:
			s.Game.fires--
			if x != x1 {
				row[x-1].RightFixed = false
			}
			if x != x2 {
				row[x+1].LeftFixed = false
			}
			ice.LeftFixed = false
			ice.RightFixed = false
			ice.TypeFlag = Blank
			down.TypeFlag = Blank
		case Blank:
			s.Game.vSwapQuite(ice, down)
		}
	}
	time.Sleep(stepTime)
	return s.Game.UpdateUI() == nil
}

func (s *Sprite) climbLeft() bool {
	return s.climb(s.LeftUp())
}

func (s *Sprite) climbRight() bool {
	return s.climb(s.RightUp())
}

func (s *Sprite) climb(dst *Sprite) bool {
	up := s.Up()
	if up != nil && up.TypeFlag != Blank {
		return false
	}
	if dst == nil || dst.TypeFlag != Blank {
		return false
	}
	return s.Game.swap(s, dst)
}

func (s *Sprite) LeftUp() *Sprite {
	left := s.Left()
	if left == nil {
		return nil
	}
	return left.Up()
}
func (s *Sprite) RightUp() *Sprite {
	right := s.Right()
	if right == nil {
		return nil
	}
	return right.Up()
}
func (s *Sprite) LeftDown() *Sprite {
	left := s.Left()
	if left == nil {
		return nil
	}
	return left.Down()
}
func (s *Sprite) RightDown() *Sprite {
	right := s.Right()
	if right == nil {
		return nil
	}
	return right.Down()
}
