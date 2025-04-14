package pkg

import (
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

var succeedMsgs = []string{
	"Wanderful!", "Brilliant!", "Excellent!", "Fantastic!", "Awesome!",
}

type Game struct {
	*amisgo.App
	levelIndex int
	levels     []Level
	PrevForm   comp.Form
	NextForm   comp.Form
	ResetForm  comp.Form
	reset      func()
	rd         *rand.Rand
	sceneName  string
	wsPath     string
	sceneFn    func() any
	wsConn     *websocket.Conn
	wsUpgrader websocket.Upgrader
}

func New(app *amisgo.App, opts ...Option) *Game {
	g := &Game{
		App: app,
		rd:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	for _, opt := range opts {
		opt(g)
	}
	g.makeLevelForms()
	g.wsUpgrader = websocket.Upgrader{}
	g.App.HandleFunc(g.wsPath, g.wsHandler)
	return g
}

func (g *Game) PrevLevel() {
	if g.levelIndex == 0 {
		return
	}
	g.levelIndex--
	g.Reset()
}

func (g *Game) NextLevel() {
	if g.levelIndex == len(g.levels)-1 {
		return
	}
	g.levelIndex++
	g.Reset()
}

func (g *Game) Reset() {
	g.reset()
}

func (g *Game) CurrentLevel() Level {
	return g.levels[g.levelIndex]
}

func (g *Game) SucceedMsg() string {
	return succeedMsgs[g.rd.Intn(len(succeedMsgs))]
}
