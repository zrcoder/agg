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
	var res []*Bar
	preX := b.Left.X
	for x := b.Left.X; x <= b.Right.X; x++ {
		cur, down := row[x], downRow[x]
		if down.Kind == Fire {
			switch cur.Kind {
			case Player:
				cur.PlayerDie()
				g.UpdateUI()
				return nil
			case Ice, IceFixed:
				cur.IceDie()
				down.FireDie()
				if preX < x {
					res = append(res, &Bar{row[preX], row[x-1]})
				}
				preX = x + 1
			}
		}
	}
	if preX <= b.Right.X {
		res = append(res, &Bar{row[preX], row[b.Right.X]})
	}
	for x := b.Left.X; x <= b.Right.X; x++ {
		g.swapQuietly(row[x], downRow[x])
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
