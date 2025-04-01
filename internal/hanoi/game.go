package hanoi

import (
	"github.com/zrcoder/agg/pkg"
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

const (
	PileCount = 3
)

type Game struct {
	*amisgo.App
	*pkg.Game
	PileA     *Pile
	PileB     *Pile
	PileC     *Pile
	piles     []*Pile
	ShiftDisk *Disk
	colors    []string
	pileForms []comp.Form
	steps     int
	codeFn    func(string) error
}

type Pile struct {
	*Game
	Index int
	Disks []*Disk
}

type Disk struct {
	*Pile
	ID int
}

func New(app *amisgo.App, codeFn func(string) error) *Game {
	g := &Game{
		App: app,
		colors: []string{
			"red", "green", "blue", "yellow", "brown", "pink", // "purple", "orange",
		},
		codeFn: codeFn,
	}
	levels := []pkg.Level{
		{Name: "Easy", Value: 2},
		{Name: "Medium", Value: 3},
		{Name: "Hard", Value: 5},
		{Name: "Expert", Value: 6},
	}
	base := pkg.New(
		app,
		pkg.WithLevels(levels, g.Reset),
		pkg.WithScene(sceneName, g.Main),
	)
	g.Game = base
	g.PileA = NewPile(g, 0)
	g.PileB = NewPile(g, 1)
	g.PileC = NewPile(g, 2)
	g.piles = []*Pile{g.PileA, g.PileB, g.PileC}
	g.makePileForms()
	g.Reset()
	return g
}

func NewPile(g *Game, index int) *Pile {
	return &Pile{
		Game:  g,
		Index: index,
		Disks: make([]*Disk, 0, len(g.colors)),
	}
}

func NewDisk(p *Pile, id int) *Disk {
	return &Disk{
		Pile: p,
		ID:   id,
	}
}

func (g *Game) Reset() {
	g.Shuffle(len(g.colors), func(i, j int) {
		g.colors[i], g.colors[j] = g.colors[j], g.colors[i]
	})
	g.PileA.renewDisks()
	g.PileB.ClearDisks()
	g.PileC.ClearDisks()
	g.ShiftDisk = nil
}

func (g *Game) IsDone() bool {
	return len(g.PileC.Disks) == g.CurrentLevel().Value
}

func (g *Game) SelectPile(pile *Pile) {
	if g.ShiftDisk != nil {
		if g.ShiftDisk.Pile == pile || pile.Empty() || pile.Top().ID > g.ShiftDisk.ID {
			pile.Push(g.ShiftDisk)
			g.ShiftDisk = nil
			return
		}
		return
	}
	g.ShiftDisk = pile.Pop()
}

func (g *Game) MinSteps() int {
	return (1 << g.CurrentLevel().Value) - 1
}

func (p *Pile) Push(d *Disk) {
	p.Disks = append(p.Disks, d)
	d.Pile = p
}

func (p *Pile) Empty() bool {
	return len(p.Disks) == 0
}

func (p *Pile) Top() *Disk {
	return p.Disks[len(p.Disks)-1]
}

func (p *Pile) Pop() *Disk {
	if p.Empty() {
		return nil
	}
	disk := p.Disks[len(p.Disks)-1]
	p.Disks = p.Disks[:len(p.Disks)-1]
	disk.Pile = nil
	return disk
}

func (p *Pile) renewDisks() {
	p.Disks = p.Disks[:0]
	for i := 0; i < p.CurrentLevel().Value; i++ {
		p.Disks = append(p.Disks, NewDisk(p, i))
	}
}

func (p *Pile) ClearDisks() {
	p.Disks = p.Disks[:0]
}
