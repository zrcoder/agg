package ballsort

import (
	_ "embed"
	"fmt"

	"github.com/zrcoder/agg/internal/ball-sort/game"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/conf"
)

//go:embed main.css
var customCSS string

func Run() {
	app := amisgo.New(
		conf.WithTitle(game.Title),
		conf.WithThemes(
			conf.Theme{Value: conf.ThemeCxd, Icon: "fa fa-sun"},
			conf.Theme{Value: conf.ThemeDark, Icon: "fa fa-moon"},
		),
		conf.WithCustomCSS(customCSS),
	)
	game := game.New(app)
	app.Mount("/", game.Index())

	fmt.Println("Server started at http://localhost:3000")
	panic(app.Run(":3000"))
}
