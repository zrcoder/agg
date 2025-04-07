package hanoi

import (
	"log"
	"net/http"
)

const (
	wsPath = "/hanoi/ws"
)

func (g *Game) wsHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	g.wsConn, err = g.wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Panic(err)
		return
	}
	defer g.wsConn.Close()

	g.updateUI()

	select {}
}

func (g *Game) updateUI() error {
	return g.wsConn.WriteJSON(map[string]any{sceneName: g.Main()})
}
