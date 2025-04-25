package vision

import (
	"github.com/RoboCup-SSL/ssl-vision-client/internal/common"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"
	"time"
)

const publishDt = 50 * time.Millisecond

func HandleVisionDetection(DetectionProvider func() *SSL_DetectionFrame) http.Handler {
	return common.UpgradeToWebsocket(
		func(r *http.Request, conn *websocket.Conn) {
			log.Println("Client for vision detection connected")
			defer log.Println("Client for vision detection disconnected")

			for {
				packet := DetectionProvider()
				if err := common.SendProtoMessage(conn, packet); err != nil {
					log.Println(err)
					return
				}

				time.Sleep(publishDt)
			}
		},
	)
}

func HandleVisionGeometry(GeometryProvider func() *SSL_GeometryData) http.Handler {
	return common.UpgradeToWebsocket(
		func(r *http.Request, conn *websocket.Conn) {
			log.Println("Client for vision geometry connected")
			defer log.Println("Client for vision geometry disconnected")

			var lastPacket *SSL_GeometryData
			for {
				packet := GeometryProvider()
				if lastPacket == nil || !proto.Equal(packet, lastPacket) {
					if err := common.SendProtoMessage(conn, packet); err != nil {
						log.Println(err)
						return
					}
					lastPacket = packet
				}

				time.Sleep(publishDt)
			}
		},
	)
}
