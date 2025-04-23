package tracked

import (
	"github.com/RoboCup-SSL/ssl-vision-client/internal/common"
	"github.com/gorilla/websocket"
	"log"
	"maps"
	"net/http"
	"time"
)

const publishDt = 50 * time.Millisecond

type TrackerSourceResponse struct {
	TrackerSources map[string]string `json:"tracker_sources"`
}

type TrackerSourceRequest struct {
	TrackerSource string `json:"tracker_source"`
}

func HandleTrackerSources(TrackerProvider func() map[string]*TrackerWrapperPacket) http.Handler {
	return common.UpgradeToWebsocket(
		func(r *http.Request, conn *websocket.Conn) {
			var trackerSourceMap map[string]string
			for {
				packet := TrackerProvider()
				newTrackerSourceMap := getTrackerSourceMap(packet)
				if !maps.Equal(newTrackerSourceMap, trackerSourceMap) {
					trackerSourceMap = newTrackerSourceMap
					message := TrackerSourceResponse{TrackerSources: trackerSourceMap}

					if err := common.SendJSONMessage(conn, message); err != nil {
						log.Println(err)
						return
					}
				}

				time.Sleep(time.Millisecond * 300)
			}
		},
	)
}

func getTrackerSourceMap(packet map[string]*TrackerWrapperPacket) map[string]string {
	trackerSourceMap := map[string]string{}
	for source, frame := range packet {
		sourceName := frame.SourceName
		if sourceName != nil {
			trackerSourceMap[source] = *sourceName
		} else {
			trackerSourceMap[source] = "Unknown"
		}
	}
	return trackerSourceMap
}

func HandleTracker(TrackerProvider func() map[string]*TrackerWrapperPacket) http.Handler {
	return common.UpgradeToWebsocket(
		func(r *http.Request, conn *websocket.Conn) {
			activeTrackerSource := r.URL.Query().Get("source")
			log.Printf("Client for tracker source '%v' connected", activeTrackerSource)
			defer log.Printf("Client for tracker source '%v' disconnected", activeTrackerSource)

			for {
				packets := TrackerProvider()
				if packet, ok := packets[activeTrackerSource]; ok {
					if err := common.SendProtoMessage(conn, packet.TrackedFrame); err != nil {
						log.Println(err)
						return
					}
				}

				time.Sleep(publishDt)
			}
		},
	)
}
