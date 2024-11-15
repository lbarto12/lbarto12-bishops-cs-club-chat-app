package handlers

import (
	"api/postgres"
	"encoding/json"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"strconv"
)

func handleConnection(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		has, err := strconv.ParseInt(string(p), 10, 64)
		if err != nil {
			log.Println(err)
			return
		}

		var msgs uint64 = uint64(has)
		for uint64(has) >= msgs {
			msgs, err = postgres.GetNumMessages()
			if err != nil {
				log.Println(err)
			}
		}

		nw, err := postgres.GetRecent(msgs - uint64(has))
		if err != nil {
			log.Println(err)
			return
		}

		marshalled, err := json.Marshal(nw)
		if err != nil {
			log.Println(err)
			return
		}

		if err = conn.WriteMessage(websocket.TextMessage, marshalled); err != nil {
			log.Println(err)
			return
		}
	}
}

func AddLiveHandler(mux *http.ServeMux) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	mux.HandleFunc("/live", func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		handleConnection(conn)
	})
}
