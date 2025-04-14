package hanoi

import (
	"time"

	"github.com/zrcoder/agg/internal"
	"github.com/zrcoder/agg/internal/hanoi"
)

func A() {
	tap(internal.Agg.Hanoi.PileA)
}

func B() {
	tap(internal.Agg.Hanoi.PileB)
}

func C() {
	tap(internal.Agg.Hanoi.PileC)
}

func tap(p *hanoi.Pile) {
	internal.Agg.Hanoi.SelectPile(p)
	time.Sleep(500 * time.Millisecond)
	internal.Agg.Hanoi.UpdateUI()
}
