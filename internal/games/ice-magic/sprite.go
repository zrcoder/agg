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
	Ice      = 'i'
	IceFixed = 'I'
)

var (
	stepTime = 100 * time.Millisecond
)

type Sprite struct {
	*Game
	Kind       byte
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
	td := s.Td().Colspan(1).Width("40px").Align("center")
	switch s.Kind {
	case Fire:
		td.Body(s.Tpl("🔥")).Style(noBorderTdStyle)
	case Player:
		td.Body(s.Tpl("🐙")).Style(noBorderTdStyle)
	case Wall:
		td.Background("#A52A2A").Style(s.borderStyle())
	case Ice:
		td.Background("#87CEFA").Style(s.borderStyle())
	case Blank:
		td.Style(noBorderTdStyle)
	}
	return td
}

func (s *Sprite) Tpl(text string) comp.Tpl {
	return s.Game.Tpl().Tpl(text).ClassName("text-2xl")
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
	switch s.Kind {
	case Player:
		return s.playerMoveLeft()
	case Ice:
		return s.iceSlideLeft()
	}
	return false
}

func (s *Sprite) moveRight() bool {
	switch s.Kind {
	case Player:
		return s.playerMoveRight()
	case Ice:
		return s.iceSlideRight()
	}
	return false
}

func (s *Sprite) iceSlideLeft() bool {
	left := s.Left()
	if left == nil {
		return false
	}
	switch left.Kind {
	case Fire:
		left.FireDie()
		s.IceDie()
		time.Sleep(stepTime)
		s.Game.UpdateUI()
		s.Game.checkFall(left.Up())
		s.Game.checkFall(s.Up())
		return true
	case Blank:
		up := s.Up()
		if ok := s.Game.swap(left, s, stepTime); !ok {
			return false
		}
		if ok := s.fall(); !ok {
			s.iceSlideLeft()
		}
		s.Game.checkFall(up)
		return true
	}
	return false
}
func (s *Sprite) iceSlideRight() bool {
	right := s.Right()
	if right == nil {
		return false
	}
	switch right.Kind {
	case Fire:
		right.FireDie()
		s.IceDie()
		time.Sleep(stepTime)
		s.Game.UpdateUI()
		s.Game.checkFall(right.Up())
		s.Game.checkFall(s.Up())
		return true
	case Blank:
		up := s.Up()
		if ok := s.Game.swap(s, right, stepTime); !ok {
			return false
		}
		if ok := s.fall(); !ok {
			s.iceSlideRight()
		}
		s.Game.checkFall(up)
		return true
	}
	return false
}

func (s *Sprite) fall1step() bool {
	if s == nil || s.Y >= len(s.Game.grid) {
		return false
	}
	down := s.Down()
	switch s.Kind {
	case Player:
		switch down.Kind {
		case Blank:
			return s.Game.swapQuietly(s, down)
		case Fire:
			s.PlayerDie()
			down.FireDie()
			return true
		}
	case Ice:
		switch down.Kind {
		case Blank:
			return s.Game.swapQuietly(s, down)
		case Fire:
			s.IceDie()
			down.FireDie()
			return true
		}
	case Fire:
		if down.Kind == Blank {
			return s.Game.swapQuietly(s, down)
		}
	}
	return false
}

func (s *Sprite) climbLeft() bool {
	return s.climb(s.LeftUp())
}

func (s *Sprite) climbRight() bool {
	return s.climb(s.RightUp())
}

func (s *Sprite) climb(dst *Sprite) bool {
	up := s.Up()
	if up != nil && up.Kind != Blank {
		return false
	}
	if dst == nil || dst.Kind != Blank {
		return false
	}
	return s.Game.swap(s, dst, 2*stepTime)
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

func (s *Sprite) IsIce() bool {
	return s.Kind == Ice || s.Kind == IceFixed
}

func (s *Sprite) IceDie() {
	if !s.IsIce() {
		return
	}
	s.Kind = Blank
	s.UnFix()
}
func (s *Sprite) FireDie() {
	if s.Kind != Fire {
		return
	}
	s.Game.fires--
	s.Kind = Blank
}
func (s *Sprite) PlayerDie() {
	if s.Kind != Player {
		return
	}
	s.Game.failed = true
	s.Kind = Blank
}
func (s *Sprite) UnFix() {
	s.LeftFixed = false
	s.RightFixed = false
	left, right := s.Left(), s.Right()
	if left != nil {
		left.RightFixed = false
	}
	if right != nil {
		right.LeftFixed = false
	}
}

func (s *Sprite) getIceBar() *Bar {
	if s == nil || !s.IsIce() {
		return nil
	}
	x1, x2 := s.X, s.X
	row := s.Game.grid[s.Y]
	for x1 >= 0 && row[x1].IsIce() && row[x1].LeftFixed {
		x1--
	}
	if !row[x1].IsIce() {
		x1++
	}
	for x2 < len(row) && row[x2].IsIce() && row[x2].RightFixed {
		x2++
	}
	if !row[x2].IsIce() {
		x2--
	}
	return &Bar{Left: row[x1], Right: row[x2]}
}

func (s *Sprite) magicLeft() {
	if s.Kind != Player {
		return
	}
	s.magic(s.LeftDown())

}
func (s *Sprite) magicRight() {
	if s.Kind != Player {
		return
	}
	s.magic((s.RightDown()))
}

func (s *Sprite) magic(dst *Sprite) {
	if dst == nil {
		return
	}
	switch dst.Kind {
	case Blank:
		dst.Kind = Ice
		left := dst.Left()
		right := dst.Right()
		if left != nil && (left.Kind == Ice || left.Kind == Wall) {
			left.RightFixed = true
			dst.LeftFixed = true
		}
		if right != nil && (right.Kind == Ice || right.Kind == Wall) {
			right.LeftFixed = true
			dst.RightFixed = true
		}
		s.Game.UpdateUI()
	case Ice:
		up, left, right := dst.Up(), dst.Left(), dst.Right()
		dst.IceDie()
		s.Game.UpdateUI()
		s.Game.checkFall(up)
		s.Game.checkFall(left)
		s.Game.checkFall(right)
	}
}

func (s *Sprite) fall() bool {
	return s.Game.FallBars(&Bar{s, s})
}
