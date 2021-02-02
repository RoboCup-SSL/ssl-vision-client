package vision

import (
	"github.com/RoboCup-SSL/ssl-vision-client/pkg/sslnet"
	"github.com/golang/protobuf/proto"
	"log"
	"sync"
	"time"
)

type Receiver struct {
	detections        map[int]*SSL_DetectionFrame
	receivedTimes     map[int]time.Time
	Geometry          *SSL_GeometryData
	mutex             sync.Mutex
	MulticastReceiver *sslnet.MulticastReceiver
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
	r.MulticastReceiver = sslnet.NewMulticastReceiver(r.consumeMessage)
	return
}

func (r *Receiver) Start(multicastAddress string) {
	r.MulticastReceiver.Start(multicastAddress)
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

func (r *Receiver) consumeMessage(data []byte) {
	message, err := parseVisionWrapperPacket(data)
	if err != nil {
		log.Print("Could not parse message: ", err)
		return
	}
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
