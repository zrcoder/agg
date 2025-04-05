package hanoi

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	wsPath = "/hanoi/ws"
)

var (
	upgrader = websocket.Upgrader{}
	wsConn   *websocket.Conn
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	wsConn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer wsConn.Close()

	select {}
}

func updateUI(ui any) error {
	data := map[string]any{
		"main": ui,
	}
	return wsConn.WriteJSON(data)
}
