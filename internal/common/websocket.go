package common

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"
)

func UpgradeToWebsocket(handleConnection func(*http.Request, *websocket.Conn)) http.Handler {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(*http.Request) bool { return true },
	}

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			conn, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				log.Println("Could not upgrade connection: ", err)
				return
			}

			defer func() {
				if err := conn.Close(); err != nil {
					log.Println("Could not close websocket connection: ", err)
				}
			}()

			handleConnection(r, conn)
		},
	)
}

func SendProtoMessage(conn *websocket.Conn, message proto.Message) error {
	if payload, err := proto.Marshal(message); err != nil {
		return errors.Wrap(err, "Failed to marshall message to protobuf")
	} else {
		return sendMessage(conn, websocket.BinaryMessage, payload)
	}
}

func SendJSONMessage(conn *websocket.Conn, message interface{}) error {
	if payload, err := json.Marshal(message); err != nil {
		return errors.Wrap(err, "Failed to marshall message to JSON")
	} else {
		return sendMessage(conn, websocket.TextMessage, payload)
	}
}

func sendMessage(conn *websocket.Conn, messageType int, message []byte) error {
	if err := conn.WriteMessage(messageType, message); err != nil {
		return errors.Wrap(err, "Failed to send message")
	}
	return nil
}
