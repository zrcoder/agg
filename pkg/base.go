package pkg

import (
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/schema"
)

type Base struct {
	*amisgo.App
	chapters       []Chapter
	chapterOptions []any
	chapterIndex   int
	levels         []Level
	levelOptions   []any
	levelIndex     int
	ResetForm      comp.Form
	LeveSelectForm comp.Form
	reset          func()
	sceneName      string
	rd             *rand.Rand
	successMsgs    []string
	wsPath         string
	sceneFn        func() any
	wsConn         *websocket.Conn
	wsUpgrader     websocket.Upgrader
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
	b.ResetForm = b.resetForm()
	b.LeveSelectForm = b.levelSelectForm()
}

func (b *Base) resetForm() comp.Form {
	return b.Form().Mode("inline").WrapWithPanel(false).Submit(
		func(s schema.Schema) error {
			b.reset()
			return b.UpdateUI()
		}).
		Body(
			b.Button().ActionType("submit").Label("Ctrl+R").Icon("fa fa-refresh").HotKey("ctrl+r"),
		)
}

func (b *Base) levelSelectForm() comp.Form {
	var options []any
	var value string
	if len(b.chapters) > 0 {
		options = b.chapterOptions
		value = makeChapterLevelOptionValue(b.chapterIndex, b.levelIndex)
	} else {
		options = b.levelOptions
		value = b.levels[b.LevelIndex()].Value
	}
	const levelSelectID = "levelSelect"
	return b.Form().Mode("inline").WrapWithPanel(false).SubmitOnChange(true).Submit(
		func(s schema.Schema) error {
			index := s.Get(levelSelectID).(string)
			if len(b.chapters) == 0 {
				b.levelIndex, _ = strconv.Atoi(index)
			} else {
				b.chapterIndex, b.levelIndex = calChapterLevelIndex(index)
			}
			b.reset()
			return b.UpdateUI()
		}).
		Body(
			b.Select().Name(levelSelectID).Label("LEVEL").SelectMode("chained").LabelClassName("text-xl font-bold").Value(value).Options(
				options...,
			),
		)
}
