package gc

import (
	"github.com/RoboCup-SSL/ssl-vision-client/internal/common"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const publishDt = 100 * time.Millisecond

func HandleReferee(RefereeProvider func() *Referee) http.Handler {
	return common.UpgradeToWebsocket(
		func(r *http.Request, conn *websocket.Conn) {
			log.Println("Client for referee connected")
			for {
				packet := RefereeProvider()
				if packet != nil {
					if err := common.SendProtoMessage(conn, packet); err != nil {
						log.Println(err)
						return
					}
				}

				time.Sleep(publishDt)
			}
		},
	)
}
