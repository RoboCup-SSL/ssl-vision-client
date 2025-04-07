package tracked

import (
	"github.com/RoboCup-SSL/ssl-vision-client/internal/common"
	"github.com/gorilla/websocket"
	"log"
	"maps"
	"net/http"
	"slices"
	"sync"
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
			var trackerSources []string
			for {
				packet := TrackerProvider()
				newTrackerSources := slices.Collect(maps.Keys(packet))
				if !slices.Equal(newTrackerSources, trackerSources) {
					trackerSourceMap := getTrackerSourceMap(packet)
					message := TrackerSourceResponse{TrackerSources: trackerSourceMap}

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
			var activeTrackerSource string
			mutex := sync.Mutex{}

			go func() {
				for {
					message := TrackerSourceRequest{}
					if err := conn.ReadJSON(&message); err != nil {
						log.Println("Error reading message:", err)
						return
					} else {
						mutex.Lock()
						activeTrackerSource = message.TrackerSource
						mutex.Unlock()
						log.Printf("Client selected tracker source %v", activeTrackerSource)
					}
				}
			}()

			for {
				packets := TrackerProvider()
				mutex.Lock()
				source := activeTrackerSource
				mutex.Unlock()
				if packet, ok := packets[source]; ok {
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
