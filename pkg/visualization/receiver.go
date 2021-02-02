package visualization

import (
	"github.com/RoboCup-SSL/ssl-vision-client/pkg/sslnet"
	"github.com/golang/protobuf/proto"
	"log"
	"sync"
	"time"
)

type Receiver struct {
	frames            map[string]*VisualizationFrame
	receivedTimes     map[string]time.Time
	mutex             sync.Mutex
	multicastReceiver *sslnet.MulticastReceiver
}

func NewReceiver() (r Receiver) {
	r.frames = map[string]*VisualizationFrame{}
	r.receivedTimes = map[string]time.Time{}
	r.multicastReceiver = sslnet.NewMulticastReceiver(r.consumeMessage)
	return
}

func (r *Receiver) Start(multicastAddress string) {
	r.multicastReceiver.Start(multicastAddress)
}

func (r *Receiver) consumeMessage(data []byte) {
	frame, err := parseVisualizationFramePacket(data)
	if err != nil {
		log.Print("Could not parse referee frame: ", err)
		return
	}
	r.mutex.Lock()
	r.cleanupFrames()
	r.frames[frame.SenderId] = frame
	r.receivedTimes[frame.SenderId] = time.Now()
	r.mutex.Unlock()
}

func parseVisualizationFramePacket(data []byte) (frame *VisualizationFrame, err error) {
	frame = new(VisualizationFrame)
	err = proto.Unmarshal(data, frame)
	return
}

func (r *Receiver) cleanupFrames() {
	for id, t := range r.receivedTimes {
		if time.Now().Sub(t) > time.Second {
			delete(r.receivedTimes, id)
			delete(r.frames, id)
		}
	}
}

func (r *Receiver) GetLineSegments() map[string][]*LineSegment {
	lines := map[string][]*LineSegment{}
	r.mutex.Lock()
	for _, frame := range r.frames {
		lines[frame.SenderId] = frame.Lines
	}
	r.mutex.Unlock()
	return lines
}

func (r *Receiver) GetCircles() map[string][]*Circle {
	lines := map[string][]*Circle{}
	r.mutex.Lock()
	for _, frame := range r.frames {
		lines[frame.SenderId] = frame.Circles
	}
	r.mutex.Unlock()
	return lines
}
