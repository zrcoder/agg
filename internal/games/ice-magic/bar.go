package icemagic

type Bar struct {
	Start, End *Sprite
}

func (b *Bar) CanFall() bool {
	g := b.Start.Game
	if b.Start.Y >= len(g.grid)-1 {
		return false
	}
	down := g.grid[b.Start.Y+1]
	for x := b.Start.X; x <= b.End.X; x++ {
		switch down[x].Kind {
		case Fire, Blank:
		default:
			return false
		}
	}
	return true
}

func (b *Bar) GetUpFallBars() []*Bar {
	if b.Start.Y == 0 {
		return nil
	}
	upRow := b.Start.Game.grid[b.Start.Y-1]
	var res []*Bar
	preX := -1
	for x := b.Start.X; x <= b.End.X; x++ {
		up := upRow[x]
		switch up.Kind {
		case Wall, Blank:
		default:
			bar := up.Bar()
			if bar.Start.X > preX {
				res = append(res, bar)
				preX = bar.End.X
			}
		}
	}
	return res
}

func (b *Bar) FallBar1StepQuietly() []*Bar {
	g := b.Start.Game
	y := b.Start.Y
	if y >= len(g.grid)-1 {
		return nil
	}
	row := g.grid[y]
	downRow := g.grid[y+1]
	var res []*Bar
	preX := 0
	for x := b.Start.X; x <= b.End.X; x++ {
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
				if x-1 > preX {
					res = append(res, &Bar{downRow[preX+1], downRow[x-1]})
					preX = x
				}
			}
		}
	}
	return res
}

func (s *Sprite) Bar() *Bar {
	if s == nil {
		return nil
	}
	if s.IsIce() {
		return s.getIceBar()
	}
	return &Bar{Start: s, End: s}
}
