package referee

import (
	"github.com/RoboCup-SSL/ssl-go-tools/pkg/sslnet"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
	"sync"
)

type Receiver struct {
	lastRefereeMsg    *Referee
	mutex             sync.Mutex
	MulticastServer   *sslnet.MulticastServer
	ConsumeRefereeMsg func(msg *Referee)
}

func NewReceiver(multicastAddress string) (r *Receiver) {
	r = new(Receiver)
	r.MulticastServer = sslnet.NewMulticastServer(multicastAddress)
	r.MulticastServer.Consumer = r.consumeMessage
	r.ConsumeRefereeMsg = func(referee *Referee) {
		// noop by default
	}
	return
}

func (r *Receiver) Start() {
	r.MulticastServer.Start()
}

func (r *Receiver) RefereeMsg() (msg *Referee) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.lastRefereeMsg
}

func (r *Receiver) consumeMessage(data []byte, _ *net.UDPAddr) {
	msg := new(Referee)
	if err := proto.Unmarshal(data, msg); err != nil {
		log.Print("Could not parse message: ", err)
		return
	}
	r.mutex.Lock()
	r.lastRefereeMsg = msg
	r.mutex.Unlock()
}
