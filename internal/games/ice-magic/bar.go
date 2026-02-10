package icemagic

type Bar struct {
	Left, Right *Sprite
}

func (b *Bar) CanFall() bool {
	g := b.Left.Game
	if b.Left.Y >= len(g.grid)-1 {
		return false
	}
	if b.IceFixed() {
		return false
	}
	down := g.grid[b.Left.Y+1]
	for x := b.Left.X; x <= b.Right.X; x++ {
		switch down[x].Kind {
		case Fire, Blank:
		default:
			return false
		}
	}
	return true
}

func (b *Bar) GetUpFallBars() []*Bar {
	if b.Left.Y == 0 {
		return nil
	}
	upRow := b.Left.Game.grid[b.Left.Y-1]
	var res []*Bar
	preX := -1
	for x := b.Left.X; x <= b.Right.X; x++ {
		up := upRow[x]
		switch up.Kind {
		case Wall, Blank:
		default:
			bar := up.Bar()
			if bar.Left.X > preX {
				res = append(res, bar)
				preX = bar.Right.X
			}
		}
	}
	return res
}

func (b *Bar) FallBar1StepQuietly() []*Bar {
	g := b.Left.Game
	y := b.Left.Y
	if y >= len(g.grid)-1 {
		return nil
	}
	row := g.grid[y]
	downRow := g.grid[y+1]
	var splits []int
	for x := b.Left.X; x <= b.Right.X; x++ {
		cur, down := row[x], downRow[x]
		switch down.Kind {
		case Blank:
			g.swapQuietly(cur, down)
		case Fire:
			switch cur.Kind {
			case Player:
				cur.PlayerDie()
				g.UpdateUI()
				return nil
			case Ice, IceFixed:
				cur.IceDie()
				down.FireDie()
				splits = append(splits, x)
			}
		}
	}
	if len(splits) == 0 {
		return []*Bar{{downRow[b.Left.X], downRow[b.Right.X]}}
	}
	res := make([]*Bar, 0, len(splits)+1)
	preX := b.Left.X
	for _, x := range splits {
		if preX <= x-1 {
			res = append(res, &Bar{downRow[preX], downRow[x-1]})
		}
		preX = x + 1
	}
	if splits[len(splits)-1] != b.Right.X {
		res = append(res, &Bar{downRow[preX], downRow[b.Right.X]})
	}
	return res
}

func (b *Bar) IceFixed() bool {
	return b.Left.LeftFixed || b.Right.RightFixed
}

func (s *Sprite) Bar() *Bar {
	if s == nil {
		return nil
	}
	if s.IsIce() {
		return s.getIceBar()
	}
	return &Bar{Left: s, Right: s}
}
