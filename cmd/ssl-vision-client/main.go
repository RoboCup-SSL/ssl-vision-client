package main

import (
	"flag"
	"github.com/RoboCup-SSL/ssl-vision-client/pkg/vision"
	"github.com/gobuffalo/packr"
	"log"
	"net/http"
)

var address = flag.String("address", "localhost:8082", "The address on which the UI and API is served, default: localhost:8082")

func main() {
	visionAddress := flag.String("visionAddress", "224.5.23.2:10006", "The multicast address of ssl-vision, default: 224.5.23.2:10006")
	flag.Parse()

	setupVisionClient(*visionAddress)
	setupUi()
	err := http.ListenAndServe(*address, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func setupVisionClient(address string) {
	receiver := vision.NewReceiver()
	publisher := vision.NewPublisher()
	publisher.PackageProvider = receiver.ToPackage
	http.HandleFunc("/api/vision", publisher.Handler)
	go receiver.Receive(address)
}

func setupUi() {
	box := packr.NewBox("../../dist")
	http.Handle("/", http.FileServer(box))
	if box.Has("index.html") {
		log.Printf("UI is available at http://%v", *address)
	} else {
		log.Print("Backend-only version started. Run the UI separately or get a binary that has the UI included")
	}
}
