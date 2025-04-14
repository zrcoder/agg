package pkg

import (
	"log"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/schema"
)

type Base struct {
	*amisgo.App
	levelIndex  int
	levels      []Level
	PrevForm    comp.Form
	NextForm    comp.Form
	ResetForm   comp.Form
	reset       func()
	sceneName   string
	rd          *rand.Rand
	successMsgs []string
	wsPath      string
	sceneFn     func() any
	wsConn      *websocket.Conn
	wsUpgrader  websocket.Upgrader
}

func New(app *amisgo.App, opts ...Option) *Base {
	seed1 := uint64(time.Now().UnixNano())
	seed2 := uint64(time.Now().UnixNano())
	b := &Base{
		App:         app,
		rd:          rand.New(rand.NewPCG(seed1, seed2)),
		successMsgs: []string{"Wanderful!", "Brilliant!", "Excellent!", "Fantastic!", "Awesome!"},
	}
	for _, opt := range opts {
		opt(b)
	}
	b.wsUpgrader = websocket.Upgrader{}
	b.App.HandleFunc(b.wsPath, b.wsHandler)
	b.makeLevelForms()
	return b
}

func (b *Base) wsHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	b.wsConn, err = b.wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Panic(err)
		return
	}
	defer b.wsConn.Close()

	b.UpdateUI()

	select {}
}

func (b *Base) UpdateUI() error {
	return b.wsConn.WriteJSON(map[string]any{b.sceneName: b.sceneFn()})
}

func (b *Base) UI() comp.Service {
	return b.App.Service().Name(b.sceneName).Ws(b.wsPath).Body(
		b.Amis().Name(b.sceneName),
	)
}

func (b *Base) makeLevelForms() {
	b.PrevForm = b.levelForm(-1)
	b.NextForm = b.levelForm(1)
	b.ResetForm = b.levelForm(0)
}

func (b *Base) levelForm(delta int) comp.Form {
	var label, icon, hotkey string
	var action func()
	switch delta {
	case -1:
		hotkey = "left"
		icon = "fa fa-arrow-left"
		action = b.prevLevel
	case 1:
		hotkey = "right"
		icon = "fa fa-arrow-right"
		action = b.nextLevel
	default:
		label = "Ctrl+R"
		icon = "fa fa-refresh"
		hotkey = "ctrl+r"
		action = b.reset
	}
	return b.Form().Mode("inline").WrapWithPanel(false).Submit(
		func(s schema.Schema) error {
			action()
			return b.UpdateUI()
		}).
		Body(
			b.Button().ActionType("submit").Label(label).Icon(icon).HotKey(hotkey),
		)
}

func (b *Base) prevLevel() {
	if b.levelIndex == 0 {
		return
	}
	b.levelIndex--
	b.reset()
}

func (b *Base) nextLevel() {
	if b.levelIndex == len(b.levels)-1 {
		return
	}
	b.levelIndex++
	b.reset()
}
