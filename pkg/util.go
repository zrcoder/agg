package pkg

import (
	"github.com/zrcoder/amisgo/comp"
)

func (b *Base) Shuffle(n int, swap func(i, j int)) {
	b.rd.Shuffle(n, swap)
}

func (b *Base) CurrentLevel() Level {
	return b.levels[b.levelIndex]
}

func (b *Base) LevelUI() comp.Flex {
	return b.App.Flex().Items(
		b.PrevForm,
		b.App.Tpl().Tpl(b.CurrentLevel().Name).ClassName("text-xl font-bold text-info pr-3"),
		b.NextForm,
		b.App.Wrapper(),
		b.ResetForm,
	)
}

func (b *Base) StateUI(info string) comp.Tpl {
	return b.App.Tpl().Tpl(info).ClassName("text-xl font-bold text-info")
}

func (b *Base) SuccessUI() comp.Tpl {
	msg := b.successMsgs[b.rd.IntN(len(b.successMsgs))]
	return b.App.Tpl().Tpl(msg).ClassName("text-2xl font-bold text-success")
}

func (b *Base) DescriptionUI(description string) comp.Tpl {
	return b.App.Tpl().Tpl(description).ClassName("text-xl text-gray-500")
}

func (b *Base) Main(succeed bool, state, description string, main any) any {
	var top comp.Tpl
	if succeed {
		top = b.SuccessUI()
	} else {
		top = b.StateUI(state)
	}
	return b.App.Service().Body(
		b.App.Flex().Items(top),
		b.App.Wrapper(),
		b.App.Flex().Items(main),
		b.App.Wrapper(),
		b.App.Flex().Items(b.DescriptionUI(description)),
		b.App.Divider(),
		b.LevelUI(),
	)
}
