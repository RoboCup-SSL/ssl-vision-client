package tracked

import (
	"github.com/RoboCup-SSL/ssl-vision-client/internal/common"
	"github.com/gorilla/websocket"
	"log"
	"maps"
	"net/http"
	"slices"
	"time"
)

const publishDt = 50 * time.Millisecond

type ControlMessage struct {
	TrackerSources []string `json:"tracker_sources"`
}

func HandleTrackerControl(TrackerProvider func() map[string]*TrackerWrapperPacket) http.Handler {
	return common.UpgradeToWebsocket(
		func(r *http.Request, conn *websocket.Conn) {
			var trackerSources []string
			for {
				packet := TrackerProvider()
				newTrackerSources := slices.Collect(maps.Keys(packet))
				if !slices.Equal(newTrackerSources, trackerSources) {
					message := ControlMessage{newTrackerSources}

					if err := common.SendJSONMessage(conn, message); err != nil {
						log.Println(err)
						return
					}

					trackerSources = newTrackerSources
				}

				time.Sleep(time.Millisecond * 300)
			}
		},
	)
}

func HandleTracker(TrackerProvider func() map[string]*TrackerWrapperPacket) http.Handler {
	return common.UpgradeToWebsocket(
		func(r *http.Request, conn *websocket.Conn) {
			activeTrackerSource := r.URL.Query().Get("sourceId")
			log.Printf("Client for tracker source %v connected", activeTrackerSource)
			for {
				packets := TrackerProvider()
				if packet, ok := packets[activeTrackerSource]; ok {
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
