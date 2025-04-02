package hanoi

import (
	"fmt"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/conf"
)

var Hanoi *Game

func Run(codeFn func(string) error) {
	app := amisgo.New(
		conf.WithTitle("Tower of Hanoi"),
		conf.WithThemes(
			conf.Theme{Value: conf.ThemeDark, Icon: "fa fa-moon"},
			conf.Theme{Value: conf.ThemeAntd, Icon: "fa fa-sun"},
		),
	)
	Hanoi = New(app, codeFn)
	app.Mount("/", Hanoi.UI())

	fmt.Println("Server started at http://localhost:3000")
	panic(app.Run(":3000"))
}
