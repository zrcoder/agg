package hanoi

import (
	"fmt"

	"github.com/zrcoder/amisgo"
)

var game *Game

func Run() {
	app := amisgo.New()
	game = New(app)
	app.Mount("/", game.UI())

	fmt.Println("Server started at http://localhost:3000")
	panic(app.Run(":3000"))
}
