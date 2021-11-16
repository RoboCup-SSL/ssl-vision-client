package tracked

import (
	"github.com/RoboCup-SSL/ssl-vision-client/pkg/sslnet"
	"google.golang.org/protobuf/proto"
	"log"
	"sync"
	"time"
)

type Receiver struct {
	frames            map[string]*TrackerWrapperPacket
	receivedTimes     map[string]time.Time
	mutex             sync.Mutex
	MulticastServer   *sslnet.MulticastServer
	ConsumeDetections func(frame *TrackerWrapperPacket)
}

func NewReceiver() (r *Receiver) {
	r = new(Receiver)
	r.frames = map[string]*TrackerWrapperPacket{}
	r.receivedTimes = map[string]time.Time{}
	r.MulticastServer = sslnet.NewMulticastServer(r.consumeMessage)
	r.ConsumeDetections = func(*TrackerWrapperPacket) {}
	return
}

func (r *Receiver) Start(multicastAddress string) {
	r.MulticastServer.Start(multicastAddress)
}

func (r *Receiver) TrackedFrames() map[string]*TrackerWrapperPacket {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.cleanupOldSources()
	frames := map[string]*TrackerWrapperPacket{}
	for k, v := range r.frames {
		frames[k] = v
	}
	return frames
}

func (r *Receiver) consumeMessage(data []byte) {
	message, err := parseVisionWrapperPacket(data)
	if err != nil {
		log.Print("Could not parse message: ", err)
		return
	}
	r.mutex.Lock()
	if message.Uuid != nil && message.TrackedFrame != nil {
		r.frames[*message.Uuid] = message
		r.receivedTimes[*message.Uuid] = time.Now()
		r.ConsumeDetections(message)
	}
	r.mutex.Unlock()
}

func parseVisionWrapperPacket(data []byte) (message *TrackerWrapperPacket, err error) {
	message = new(TrackerWrapperPacket)
	err = proto.Unmarshal(data, message)
	return
}

func (r *Receiver) cleanupOldSources() {
	for uuid, t := range r.receivedTimes {
		if time.Now().Sub(t) > time.Second {
			delete(r.receivedTimes, uuid)
			delete(r.frames, uuid)
		}
	}
}
