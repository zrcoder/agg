package internal

import (
	"log/slog"
	"net/http"

	ballsort "github.com/zrcoder/agg/internal/games/ball-sort"
	"github.com/zrcoder/agg/internal/games/hanoi"
	icemagic "github.com/zrcoder/agg/internal/games/ice-magic"
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
	IceMagic        *icemagic.Game
	HanoiCodeAction func(string) error
}

var Agg *App

func Run(hanoiCodeAction func(string, func() error) error) {
	app := amisgo.New(
		conf.WithTitle(Title),
		conf.WithLocalSdk(http.FS(sdk.FS)),
		conf.WithStyles("/static/bottole-button.css"),
		conf.WithIcon("/static/agg.svg"),
	)
	Agg = &App{
		App:      app,
		Hanoi:    hanoi.New(app, hanoiCodeAction),
		BallSort: ballsort.New(app),
		IceMagic: icemagic.New(app),
	}

	app.StaticFS("/static", http.FS(static.FS))
	app.Mount("/", index())

	slog.Info("Amisgo Games started", "address", "http://localhost:3000")
	if err := app.Run(":3000"); err != nil {
		slog.Error("server error", "error", err)
	}
}

func index() comp.Page {
	app := Agg.App
	return app.Page().
		Title(app.Tpl().Tpl(Title).ClassName("text-2xl font-bold")).
		Body(
			app.Tabs().TabsMode("vertical").ClassName("border-none").Tabs(
				app.Tab().Title("Ball Sort Puzzle").Hash("ball-sort").Tab(Agg.BallSort.UI()),
				app.Tab().Title("Tower of Hanoi").Hash("hanoi").Tab(Agg.Hanoi.UI()),
				app.Tab().Title("Ice Magic").Hash("ice-magic").Tab(Agg.IceMagic.UI()),
			),
		)
}