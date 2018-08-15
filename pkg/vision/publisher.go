package vision

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const publishDt = 50 * time.Millisecond

type Publisher struct {
	upgrader        websocket.Upgrader
	PackageProvider func() *Package
}

func NewPublisher() (p Publisher) {
	p.upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(*http.Request) bool { return true },
	}
	return p
}

func (p *Publisher) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := p.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	defer log.Println("Client disconnected")

	log.Println("Client connected")

	for {
		pack := p.PackageProvider()
		payload, err := json.Marshal(*pack)
		if err != nil {
			return
		}

		if err := conn.WriteMessage(websocket.TextMessage, payload); err != nil {
			log.Println(err)
			return
		}

		time.Sleep(publishDt)
	}
}
