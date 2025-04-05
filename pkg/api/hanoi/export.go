package hanoi

import (
	"time"

	"github.com/zrcoder/agg/internal/hanoi"
)

func A() {
	tap(hanoi.Hanoi.PileA)
}

func B() {
	tap(hanoi.Hanoi.PileB)
}

func C() {
	tap(hanoi.Hanoi.PileC)
}

func tap(p *hanoi.Pile) {
	hanoi.Hanoi.SelectPile(p)
	time.Sleep(500 * time.Millisecond)
}
