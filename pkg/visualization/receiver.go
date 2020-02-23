package visualization

import (
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"sync"
	"time"
)

const maxDatagramSize = 8192

type Receiver struct {
	frames        map[string]*VisualizationFrame
	receivedTimes map[string]time.Time
	mutex         sync.Mutex
}

func NewReceiver() (r Receiver) {
	r.frames = map[string]*VisualizationFrame{}
	r.receivedTimes = map[string]time.Time{}
	return
}

func (r *Receiver) Receive(multicastAddress string) {
	listener, err := openMulticastUdpConnection(multicastAddress)
	if err != nil {
		log.Printf("Could not connect to %v: %v", multicastAddress, err)
		return
	}

	data := make([]byte, maxDatagramSize)
	for {
		n, _, err := listener.ReadFrom(data)
		if err != nil {
			log.Println("ReadFromUDP failed:", err)
			break
		}

		frame, err := parseVisualizationFramePacket(data[:n])
		if err != nil {
			log.Print("Could not parse referee frame: ", err)
			break
		} else {
			r.mutex.Lock()
			r.cleanupDetections()
			r.frames[frame.SenderId] = frame
			r.receivedTimes[frame.SenderId] = time.Now()
			r.mutex.Unlock()
		}
	}

	// wait a second and restart
	if err := listener.Close(); err != nil {
		log.Println("Could not close listener: ", err)
	}
	time.Sleep(time.Second)
	r.Receive(multicastAddress)
}

func openMulticastUdpConnection(address string) (listener *net.UDPConn, err error) {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return
	}
	listener, err = net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		return
	}
	if err := listener.SetReadBuffer(maxDatagramSize); err != nil {
		log.Println("Could not set read buffer: ", err)
	}
	log.Printf("Listening on %s", address)
	return
}

func parseVisualizationFramePacket(data []byte) (frame *VisualizationFrame, err error) {
	frame = new(VisualizationFrame)
	err = proto.Unmarshal(data, frame)
	return
}

func (r *Receiver) cleanupDetections() {
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
