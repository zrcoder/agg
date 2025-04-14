package pkg

import "github.com/zrcoder/amisgo/comp"

func (g *Game) Shuffle(n int, swap func(i, j int)) {
	g.rd.Shuffle(n, swap)
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
