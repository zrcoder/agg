package hanoi

import "time"

func A() *Pile {
	return game.PileA
}

func B() *Pile {
	return game.PileB
}

func C() *Pile {
	return game.PileC
}

func Tap(getPile func() *Pile) {
	game.SelectPile(getPile())
	time.Sleep(200 * time.Millisecond)
}
