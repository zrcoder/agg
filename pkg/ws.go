package pkg

import (
	"log"
	"net/http"
)

func (g *Game) wsHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	g.wsConn, err = g.wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Panic(err)
		return
	}
	defer g.wsConn.Close()

	g.UpdateUI()

	select {}
}

func (g *Game) UpdateUI() error {
	return g.wsConn.WriteJSON(map[string]any{g.sceneName: g.sceneFn()})
}
