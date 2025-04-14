package pkg

import (
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/schema"
)

func (g *Game) UI() comp.Service {
	return g.App.Service().Name(g.sceneName).Ws(g.wsPath).Body(
		g.Amis().Name(g.sceneName),
	)
}

func (g *Game) makeLevelForms() {
	g.PrevForm = g.levelForm(-1)
	g.NextForm = g.levelForm(1)
	g.ResetForm = g.levelForm(0)
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
			return g.UpdateUI()
		}).
		Body(
			g.Button().ActionType("submit").Label(label).Icon(icon).HotKey(hotkey),
		)
}
