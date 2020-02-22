package client

import (
	"encoding/json"
	"github.com/RoboCup-SSL/ssl-go-tools/pkg/sslproto"
	"github.com/RoboCup-SSL/ssl-vision-client/pkg/visualization"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const publishDt = 50 * time.Millisecond

type Publisher struct {
	upgrader            websocket.Upgrader
	DetectionProvider   func() *sslproto.SSL_DetectionFrame
	GeometryProvider    func() *sslproto.SSL_GeometryData
	LineSegmentProvider func() map[string][]*visualization.LineSegment
	CircleProvider      func() map[string][]*visualization.Circle
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

	for {
		pack := new(Package)
		detectionFrame := p.DetectionProvider()
		pack.AddDetectionFrame(detectionFrame)
		geometry := p.GeometryProvider()
		pack.AddGeometryShapes(geometry)

		p.addVisualization(pack)

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

func (p *Publisher) addVisualization(pack *Package) {
	allLineSegments := p.LineSegmentProvider()

	for sourceId, lineSegments := range allLineSegments {
		for _, lineSegment := range lineSegments {
			pack.AddLineSegment(sourceId, lineSegment)
		}
	}

	allCircles := p.CircleProvider()

	for sourceId, circles := range allCircles {
		for _, circle := range circles {
			pack.AddCircle(sourceId, circle)
		}
	}

	pack.SortShapes()
}
