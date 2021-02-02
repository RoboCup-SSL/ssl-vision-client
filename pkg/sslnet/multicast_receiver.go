package sslnet

import (
	"log"
	"net"
	"sync"
	"time"
)

const maxDatagramSize = 8192

type MulticastReceiver struct {
	activeIfis     map[string]bool
	consumer       func([]byte)
	mutex          sync.Mutex
	SkipInterfaces []string
}

func NewMulticastReceiver(consumer func([]byte)) (r *MulticastReceiver) {
	r = new(MulticastReceiver)
	r.activeIfis = map[string]bool{}
	r.consumer = consumer
	return
}

func (r *MulticastReceiver) Start(multicastAddress string) {
	go r.Receive(multicastAddress)
}

func (r *MulticastReceiver) Receive(multicastAddress string) {
	for {
		ifis, _ := net.Interfaces()
		for _, ifi := range ifis {
			if ifi.Flags&net.FlagMulticast == 0 || // No multicast support
				r.skipInterface(ifi.Name) {
				continue
			}
			r.mutex.Lock()
			if _, ok := r.activeIfis[ifi.Name]; !ok {
				// interface not active, (re-)start receiving
				go r.receiveOnInterface(multicastAddress, ifi)
			}
			r.mutex.Unlock()
		}
		time.Sleep(1 * time.Second)
	}
}

func (r *MulticastReceiver) skipInterface(ifiName string) bool {
	for _, skipIfi := range r.SkipInterfaces {
		if skipIfi == ifiName {
			return true
		}
	}
	return false
}

func (r *MulticastReceiver) receiveOnInterface(multicastAddress string, ifi net.Interface) {
	addr, err := net.ResolveUDPAddr("udp", multicastAddress)
	if err != nil {
		log.Printf("Could resolve multicast address %v: %v", multicastAddress, err)
		return
	}

	listener, err := net.ListenMulticastUDP("udp", &ifi, addr)
	if err != nil {
		log.Printf("Could not listen at %v: %v", multicastAddress, err)
		return
	}
	if err := listener.SetReadBuffer(maxDatagramSize); err != nil {
		log.Println("Could not set read buffer: ", err)
	}

	r.mutex.Lock()
	r.activeIfis[ifi.Name] = true
	defer delete(r.activeIfis, ifi.Name)
	r.mutex.Unlock()

	log.Printf("Listening on %s (%s)", multicastAddress, ifi.Name)

	data := make([]byte, maxDatagramSize)
	for {
		n, _, err := listener.ReadFrom(data)
		if err != nil {
			log.Println("ReadFromUDP failed:", err)
			break
		}

		r.consumer(data[:n])
	}

	log.Printf("Stop listening on %s (%s)", multicastAddress, ifi.Name)

	if err := listener.Close(); err != nil {
		log.Println("Could not close listener: ", err)
	}
}
