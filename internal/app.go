package internal

import (
	"fmt"
	"net/http"

	ballsort "github.com/zrcoder/agg/internal/ball-sort"
	"github.com/zrcoder/agg/internal/hanoi"
	"github.com/zrcoder/agg/internal/static"

	sdk "gitee.com/rdor/amis-sdk/v6"
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/conf"
)

const Title = "Amisgo Games"

type App struct {
	*amisgo.App
	Hanoi           *hanoi.Game
	BallSort        *ballsort.Game
	HanoiCodeAction func(string) error
}

var Agg *App

func Run(hanoiCodeAction func(string, func() error) error) {
	app := amisgo.New(
		conf.WithTitle(Title),
		conf.WithThemes(
			conf.Theme{Value: conf.ThemeAntd, Icon: "fa fa-sun"},
			conf.Theme{Value: conf.ThemeDark, Icon: "fa fa-moon"},
		),
		conf.WithLocalSdk(http.FS(sdk.FS)),
		conf.WithStyles("/static/bottole-button.css"),
		conf.WithIcon("/static/agg.svg"),
	)
	Agg = &App{
		App:      app,
		Hanoi:    hanoi.New(app, hanoiCodeAction),
		BallSort: ballsort.New(app),
	}

	app.StaticFS("/static", http.FS(static.FS))
	app.Mount("/", index())

	fmt.Println("Amisgo Games started at http://localhost:3000")
	panic(app.Run(":3000"))
}

func index() comp.Page {
	app := Agg.App
	return app.Page().
		Title(app.Tpl().Tpl(Title).ClassName("text-2xl font-bold")).
		Toolbar(app.ThemeButtonGroupSelect()).
		Body(
			app.Tabs().TabsMode("vertical").ClassName("border-none").Tabs(
				app.Tab().Title("Ball Sort Puzzle").Hash("ball-sort").Tab(Agg.BallSort.UI()),
				app.Tab().Title("Tower of Hanoi").Hash("hanoi").Tab(Agg.Hanoi.UI()),
			),
		)
}
