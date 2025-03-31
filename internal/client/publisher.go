package client

import (
	"encoding/json"
	"github.com/RoboCup-SSL/ssl-vision-client/internal/gc"
	"github.com/RoboCup-SSL/ssl-vision-client/internal/tracked"
	"github.com/RoboCup-SSL/ssl-vision-client/internal/vision"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

const publishDt = 50 * time.Millisecond
const visionSource = "vision"

type PublishType int

type Publisher struct {
	upgrader          websocket.Upgrader
	DetectionProvider func() *vision.SSL_DetectionFrame
	TrackerProvider   func() map[string]*tracked.TrackerWrapperPacket
	GeometryProvider  func() *vision.SSL_GeometryData
	RefereeProvider   func() *gc.Referee
}

func NewPublisher() (p Publisher) {
	p.upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(*http.Request) bool { return true },
	}
	return p
}

type PublisherClient struct {
	conn                *websocket.Conn
	activeTrackedSource string
	mutex               sync.Mutex
}

func (p *Publisher) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := p.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Could not upgrade connection: ", err)
		return
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("Could not close websocket connection: ", err)
		}
		log.Println("Client disconnected")
	}()

	log.Println("Client connected")

	client := &PublisherClient{}
	client.conn = conn
	client.activeTrackedSource = r.URL.Query().Get("sourceId")
	go client.handleClientRequests()

	for {
		pack := new(Package)
		trackedFrames := p.TrackerProvider()
		client.mutex.Lock()
		client.activeTrackedSource = selectSourceId(trackedFrames, client.activeTrackedSource)

		if client.activeTrackedSource == visionSource {
			detectionFrame := p.DetectionProvider()
			pack.AddDetectionFrame(detectionFrame)
		} else {
			if frame, ok := trackedFrames[client.activeTrackedSource]; ok {
				pack.AddTrackedFrame(frame)
			}
		}
		pack.ActiveSourceId = client.activeTrackedSource
		client.mutex.Unlock()

		geometry := p.GeometryProvider()
		pack.AddGeometryShapes(geometry)

		pack.Sources = map[string]string{}
		pack.Sources[visionSource] = visionSource
		for _, frame := range trackedFrames {
			pack.Sources[*frame.Uuid] = *frame.SourceName
		}

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

func (p *PublisherClient) handleClientRequests() {
	for {
		messageType, data, err := p.conn.ReadMessage()
		if err != nil {
			log.Println("Could not read message: ", err)
			return
		}
		if messageType == websocket.TextMessage {
			var request Request
			if err := json.Unmarshal(data, &request); err != nil {
				log.Println("Could not deserialize message: ", string(data))
			} else {
				p.mutex.Lock()
				p.activeTrackedSource = request.ActiveSourceId
				p.mutex.Unlock()
			}
		} else {
			log.Println("Got non-text message")
		}
	}
}

func selectSourceId(trackedFrames map[string]*tracked.TrackerWrapperPacket, activeSourceId string) string {
	if activeSourceId == visionSource {
		return activeSourceId
	}
	if _, ok := trackedFrames[activeSourceId]; ok {
		return activeSourceId
	}
	for k := range trackedFrames {
		return k
	}
	return visionSource
}
