package pkg

import (
	"math/rand"
	"time"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/schema"
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
	sceneFn    func() any
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
	return g
}

func (g *Game) PrevLevel() {
	g.levelIndex = (g.levelIndex - 1 + len(g.levels)) % len(g.levels)
	g.Reset()
}

func (g *Game) NextLevel() {
	g.levelIndex = (g.levelIndex + 1) % len(g.levels)
	g.Reset()
}

func (g *Game) Reset() {
	g.reset()
}

func (g *Game) CurrentLevel() Level {
	return g.levels[g.levelIndex]
}

func (g *Game) makeLevelForms() {
	g.PrevForm = g.levelForm(-1)
	g.NextForm = g.levelForm(1)
	g.ResetForm = g.levelForm(0)
}

func (g *Game) LevelUI() comp.Flex {
	return g.App.Flex().Items(
		g.PrevForm,
		g.App.Tpl().Tpl(g.CurrentLevel().Name).ClassName("text-xl font-bold text-info pr-3"),
		g.NextForm,
		g.App.Wrapper(),
		g.ResetForm,
	)
}

func (g *Game) levelForm(delta int) comp.Form {
	var label, icon, hotkey string
	var action func()
	switch delta {
	case -1:
		hotkey = "left"
		icon = "fa fa-arrow-left"
		action = g.PrevLevel
	case 1:
		hotkey = "right"
		icon = "fa fa-arrow-right"
		action = g.NextLevel
	default:
		label = "Ctrl+R"
		icon = "fa fa-refresh"
		hotkey = "ctrl+r"
		action = g.Reset
	}
	return g.Form().Mode("inline").WrapWithPanel(false).Submit(
		func(s schema.Schema) error {
			action()
			return nil
		}).
		Body(
			g.Button().ActionType("submit").Label(label).Icon(icon).HotKey(hotkey).
				OnEvent(g.Event().Click(
					g.EventActions(
						g.EventAction().ActionType("refresh"),
					),
				)),
		)
}

func (g *Game) SucceedMsg() string {
	return succeedMsgs[g.rd.Intn(len(succeedMsgs))]
}

func (g *Game) StateUI(info string) comp.Tpl {
	infoClass := "text-xl font-bold text-info"
	return g.App.Tpl().Tpl(info).ClassName(infoClass)
}

func (g *Game) SucceedUI() comp.Tpl {
	return g.App.Tpl().Tpl(g.SucceedMsg()).ClassName("text-2xl font-bold text-success")
}

func (g *Game) DescriptionUI(description string) comp.Tpl {
	return g.App.Tpl().Tpl(description).ClassName("text-xl text-gray-500")
}

func (g *Game) Main(succeed bool, state, description string, main any) any {
	var top comp.Tpl
	if succeed {
		top = g.SucceedUI()
	} else {
		top = g.StateUI(state)
	}
	return g.App.Service().Body(
		g.App.Flex().Items(top),
		g.App.Wrapper(),
		g.App.Flex().Items(main),
		g.App.Wrapper(),
		g.App.Flex().Items(g.DescriptionUI(description)),
		g.App.Divider(),
		g.LevelUI(),
	)
}

func (g *Game) Service() comp.Service {
	return g.App.Service().Name(g.sceneName).GetSchema(g.sceneFn).Messages(schema.Schema{}).SilentPolling(false)
}
