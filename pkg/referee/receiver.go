package referee

import (
	"github.com/RoboCup-SSL/ssl-vision-client/pkg/sslnet"
	"google.golang.org/protobuf/proto"
	"log"
	"sync"
)

type Receiver struct {
	lastRefereeMsg    *Referee
	mutex             sync.Mutex
	MulticastServer   *sslnet.MulticastServer
	ConsumeRefereeMsg func(msg *Referee)
}

func NewReceiver() (r *Receiver) {
	r = new(Receiver)
	r.MulticastServer = sslnet.NewMulticastServer(r.consumeMessage)
	r.ConsumeRefereeMsg = func(referee *Referee) {}
	return
}

func (r *Receiver) Start(multicastAddress string) {
	r.MulticastServer.Start(multicastAddress)
}

func (r *Receiver) RefereeMsg() (msg *Referee) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.lastRefereeMsg
}

func (r *Receiver) consumeMessage(data []byte) {
	msg := new(Referee)
	if err := proto.Unmarshal(data, msg); err != nil {
		log.Print("Could not parse message: ", err)
		return
	}
	r.mutex.Lock()
	r.lastRefereeMsg = msg
	r.mutex.Unlock()
}
