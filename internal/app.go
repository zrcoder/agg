package internal

import (
	_ "embed"
	"fmt"
	"net/http"

	sdk "gitee.com/rdor/amis-sdk/v6"
	ballsort "github.com/zrcoder/agg/internal/ball-sort"
	"github.com/zrcoder/agg/internal/hanoi"
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/conf"
)

//go:embed bottole-button.css
var customCSS string

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
			conf.Theme{Value: conf.ThemeDark, Icon: "fa fa-moon"},
			conf.Theme{Value: conf.ThemeAntd, Icon: "fa fa-sun"},
		),
		conf.WithLocalSdk(http.FS(sdk.FS)),
		conf.WithCustomCSS(customCSS),
	)
	Agg = &App{
		App:      app,
		Hanoi:    hanoi.New(app, hanoiCodeAction),
		BallSort: ballsort.New(app),
	}

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
				app.Tab().Title("Ball Sort Puzzle").Tab(Agg.BallSort.UI()),
				app.Tab().Title("Tower of Hanoi").Tab(Agg.Hanoi.UI()),
			),
		)
}
