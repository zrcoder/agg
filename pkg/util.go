package pkg

func (g *Game) Shuffle(n int, swap func(i, j int)) {
	g.rd.Shuffle(n, swap)
}
