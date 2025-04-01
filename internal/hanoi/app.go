package hanoi

import (
	"fmt"

	"github.com/zrcoder/amisgo"
)

var Hanoi *Game

func Run(codeFn func(string) error) {
	app := amisgo.New()
	Hanoi = New(app, codeFn)
	app.Mount("/", Hanoi.UI())

	fmt.Println("Server started at http://localhost:3000")
	panic(app.Run(":3000"))
}
