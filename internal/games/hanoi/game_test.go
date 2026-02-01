package hanoi

import (
	"testing"

	"github.com/zrcoder/amisgo"
)

func TestNew(t *testing.T) {
	app := amisgo.New()
	codeAction := func(s string, f func() error) error { return nil }

	game := New(app, codeAction)

	if game == nil {
		t.Fatal("New() returned nil")
	}
	if game.App == nil {
		t.Error("New() did not set App")
	}
	if game.Base == nil {
		t.Error("New() did not set Base")
	}
	if game.PileA == nil {
		t.Error("New() did not set PileA")
	}
	if game.PileB == nil {
		t.Error("New() did not set PileB")
	}
	if game.PileC == nil {
		t.Error("New() did not set PileC")
	}
	if len(game.colors) != 6 {
		t.Errorf("New() did not set colors, got %v", len(game.colors))
	}
}

func TestNewPile(t *testing.T) {
	app := amisgo.New()
	game := &Game{App: app, colors: make([]string, 6)}

	pile := NewPile(game, 1)

	if pile == nil {
		t.Fatal("NewPile() returned nil")
	}
	if pile.Game != game {
		t.Error("NewPile() did not set Game")
	}
	if pile.Index != 1 {
		t.Errorf("NewPile() set Index = %v, want 1", pile.Index)
	}
	if pile.Disks == nil {
		t.Error("NewPile() did not initialize Disks")
	}
	if cap(pile.Disks) != 6 {
		t.Errorf("NewPile() set Disks capacity = %v, want 6", cap(pile.Disks))
	}
}

func TestNewDisk(t *testing.T) {
	app := amisgo.New()
	game := &Game{App: app}
	pile := NewPile(game, 0)

	disk := NewDisk(pile, 3)

	if disk == nil {
		t.Fatal("NewDisk() returned nil")
	}
	if disk.Pile != pile {
		t.Error("NewDisk() did not set Pile")
	}
	if disk.ID != 3 {
		t.Errorf("NewDisk() set ID = %v, want 3", disk.ID)
	}
}

func TestReset(t *testing.T) {
	app := amisgo.New()
	game := New(app, func(s string, f func() error) error { return nil })

	game.steps = 10
	game.PileA.Push(NewDisk(game.PileA, 1))
	game.PileB.Push(NewDisk(game.PileB, 2))
	game.ShiftDisk = NewDisk(game.PileC, 3)

	game.Reset()

	if game.steps != 0 {
		t.Errorf("Reset() did not reset steps, got %v", game.steps)
	}
	if game.ShiftDisk != nil {
		t.Error("Reset() did not reset ShiftDisk")
	}
	if len(game.PileB.Disks) != 0 {
		t.Errorf("Reset() did not clear PileB, got %v disks", len(game.PileB.Disks))
	}
	if len(game.PileC.Disks) != 0 {
		t.Errorf("Reset() did not clear PileC, got %v disks", len(game.PileC.Disks))
	}
}

func TestPilePush(t *testing.T) {
	app := amisgo.New()
	game := &Game{App: app}
	pile := NewPile(game, 0)

	disk := NewDisk(pile, 1)
	pile.Push(disk)

	if len(pile.Disks) != 1 {
		t.Errorf("Push() did not add disk, got %v", len(pile.Disks))
	}
	if pile.Disks[0] != disk {
		t.Error("Push() did not add the correct disk")
	}
	if disk.Pile != pile {
		t.Error("Push() did not update disk's Pile")
	}
}

func TestPilePop(t *testing.T) {
	app := amisgo.New()
	game := &Game{App: app}
	pile := NewPile(game, 0)

	disk1 := NewDisk(pile, 1)
	disk2 := NewDisk(pile, 2)
	pile.Push(disk1)
	pile.Push(disk2)

	popped := pile.Pop()

	if popped != disk2 {
		t.Errorf("Pop() returned wrong disk, got %v", popped)
	}
	if len(pile.Disks) != 1 {
		t.Errorf("Pop() did not remove disk, got %v", len(pile.Disks))
	}
	if pile.Disks[0] != disk1 {
		t.Error("Pop() removed the wrong disk")
	}

	popped = pile.Pop()
	if popped != disk1 {
		t.Errorf("Pop() returned wrong disk, got %v", popped)
	}
	if len(pile.Disks) != 0 {
		t.Errorf("Pop() did not remove disk, got %v", len(pile.Disks))
	}

	popped = pile.Pop()
	if popped != nil {
		t.Errorf("Pop() from empty pile returned %v, want nil", popped)
	}
}

func TestPileTop(t *testing.T) {
	app := amisgo.New()
	game := &Game{App: app}
	pile := NewPile(game, 0)

	if !pile.Empty() {
		t.Error("Empty() returned false for empty pile")
	}

	disk1 := NewDisk(pile, 1)
	disk2 := NewDisk(pile, 2)
	pile.Push(disk1)
	pile.Push(disk2)

	top := pile.Top()
	if top != disk2 {
		t.Errorf("Top() returned wrong disk, got %v", top)
	}
}

func TestPileEmpty(t *testing.T) {
	app := amisgo.New()
	game := &Game{App: app}
	pile := NewPile(game, 0)

	if !pile.Empty() {
		t.Error("Empty() returned false for empty pile")
	}

	pile.Push(NewDisk(pile, 1))
	if pile.Empty() {
		t.Error("Empty() returned true for non-empty pile")
	}
}

func TestSelectPile(t *testing.T) {
	app := amisgo.New()
	game := New(app, func(s string, f func() error) error { return nil })
	game.Reset()

	err := game.SelectPile(game.PileA)
	if err != nil {
		t.Errorf("SelectPile() returned error: %v", err)
	}
	if game.ShiftDisk == nil {
		t.Error("SelectPile() did not set ShiftDisk")
	}
	if game.ShiftDisk.ID != 0 {
		t.Errorf("SelectPile() got wrong ShiftDisk ID = %v", game.ShiftDisk.ID)
	}

	err = game.SelectPile(game.PileB)
	if err != nil {
		t.Errorf("SelectPile() returned error: %v", err)
	}
	if game.ShiftDisk != nil {
		t.Error("SelectPile() did not move disk to PileB")
	}
	if len(game.PileB.Disks) != 1 {
		t.Errorf("SelectPile() did not add disk to PileB, got %v", len(game.PileB.Disks))
	}
}

func TestMinSteps(t *testing.T) {
	app := amisgo.New()
	game := New(app, func(s string, f func() error) error { return nil })

	expected := (1 << game.CurrentLevel().Disks) - 1
	if game.MinSteps() != expected {
		t.Errorf("MinSteps() = %v, want %v", game.MinSteps(), expected)
	}
}

func TestIsDone(t *testing.T) {
	app := amisgo.New()
	game := New(app, func(s string, f func() error) error { return nil })
	game.Reset()

	if game.IsDone() {
		t.Error("IsDone() returned true at start")
	}

	for i := 0; i < game.CurrentLevel().Disks; i++ {
		game.PileC.Push(NewDisk(game.PileC, i))
	}

	if !game.IsDone() {
		t.Error("IsDone() returned false when PileC has all disks")
	}
}
