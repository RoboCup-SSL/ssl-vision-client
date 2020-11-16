package vision

import (
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"sync"
	"time"
)

const maxDatagramSize = 8192

type Receiver struct {
	detections    map[int]*SSL_DetectionFrame
	receivedTimes map[int]time.Time
	Geometry      *SSL_GeometryData
	mutex         sync.Mutex
}

func NewReceiver() (r Receiver) {
	r.detections = map[int]*SSL_DetectionFrame{}
	r.receivedTimes = map[int]time.Time{}
	r.Geometry = new(SSL_GeometryData)
	r.Geometry.Field = new(SSL_GeometryFieldSize)
	r.Geometry.Field.FieldWidth = new(int32)
	r.Geometry.Field.FieldLength = new(int32)
	r.Geometry.Field.GoalDepth = new(int32)
	r.Geometry.Field.GoalWidth = new(int32)
	r.Geometry.Field.BoundaryWidth = new(int32)
	*r.Geometry.Field.FieldWidth = 9000
	*r.Geometry.Field.FieldLength = 12000
	*r.Geometry.Field.GoalDepth = 180
	*r.Geometry.Field.GoalWidth = 1000
	*r.Geometry.Field.BoundaryWidth = 300
	return
}

func (r *Receiver) Detections() (result map[int]SSL_DetectionFrame) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	result = map[int]SSL_DetectionFrame{}
	for id, frame := range r.detections {
		result[id] = *frame
	}
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

		message, err := parseVisionWrapperPacket(data[:n])
		if err != nil {
			log.Print("Could not parse message: ", err)
			break
		} else {
			if message.Detection != nil {
				r.mutex.Lock()
				camId := int(*message.Detection.CameraId)
				r.detections[camId] = message.Detection
				r.receivedTimes[camId] = time.Now()
				r.mutex.Unlock()
			}
			if message.Geometry != nil {
				r.Geometry = message.Geometry
			}
		}
	}

	// wait a second and restart
	if err := listener.Close(); err != nil {
		log.Println("Could not close listener: ", err)
	}
	time.Sleep(time.Second)
	r.Receive(multicastAddress)
}

func (r *Receiver) CombinedDetectionFrames() (f *SSL_DetectionFrame) {
	r.mutex.Lock()
	f = new(SSL_DetectionFrame)
	f.Balls = make([]*SSL_DetectionBall, 0)
	f.RobotsYellow = make([]*SSL_DetectionRobot, 0)
	f.RobotsBlue = make([]*SSL_DetectionRobot, 0)

	r.cleanupDetections()
	for _, b := range r.detections {
		f.Balls = append(f.Balls, b.Balls...)
		f.RobotsYellow = append(f.RobotsYellow, b.RobotsYellow...)
		f.RobotsBlue = append(f.RobotsBlue, b.RobotsBlue...)
	}

	r.mutex.Unlock()
	return
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

func parseVisionWrapperPacket(data []byte) (message *SSL_WrapperPacket, err error) {
	message = new(SSL_WrapperPacket)
	err = proto.Unmarshal(data, message)
	return
}

func (r *Receiver) cleanupDetections() {
	for camId, t := range r.receivedTimes {
		if time.Now().Sub(t) > time.Second {
			delete(r.receivedTimes, camId)
			delete(r.detections, camId)
		}
	}
}

func (r *Receiver) CurrentGeometry() (geometry *SSL_GeometryData) {
	geometry = r.Geometry
	return
}
