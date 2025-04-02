package hanoi

import (
	"time"

	"github.com/zrcoder/agg/internal/hanoi"
)

type Pile = func() *hanoi.Pile

func A() *hanoi.Pile {
	return hanoi.Hanoi.PileA
}

func B() *hanoi.Pile {
	return hanoi.Hanoi.PileB
}

func C() *hanoi.Pile {
	return hanoi.Hanoi.PileC
}

func Tap(getPile Pile) {
	hanoi.Hanoi.SelectPile(getPile())
	time.Sleep(200 * time.Millisecond)
}
