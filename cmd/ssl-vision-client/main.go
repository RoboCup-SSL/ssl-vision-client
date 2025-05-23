package main

import (
	"flag"
	"github.com/RoboCup-SSL/ssl-vision-client/internal/gc"
	"github.com/RoboCup-SSL/ssl-vision-client/internal/tracked"
	"github.com/RoboCup-SSL/ssl-vision-client/internal/vision"
	"log"
	"net/http"
	"strings"
)

var address = flag.String("address", ":8082", "The address on which the UI and API is served, default: :8082")
var visionAddress = flag.String("visionAddress", "224.5.23.2:10006", "The multicast address of ssl-vision, default: 224.5.23.2:10006")
var trackedAddress = flag.String("trackedAddress", "224.5.23.2:10010", "The multicast address of trackers, default: 224.5.23.2:10010")
var refereeAddress = flag.String("refereeAddress", "224.5.23.1:10003", "The multicast address of the game controller, default: 224.5.23.1:10003")
var skipInterfaces = flag.String("skipInterfaces", "", "Comma separated list of interface names to ignore when receiving multicast packets")
var verbose = flag.Bool("verbose", false, "Verbose output")

func main() {
	flag.Parse()

	log.Printf("UI is available at %v", formattedAddress())

	httpServer := setupServer()

	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func setupServer() *http.Server {
	visionReceiver := vision.NewReceiver(*visionAddress)
	trackedReceiver := tracked.NewReceiver(*trackedAddress)
	refereeReceiver := gc.NewReceiver(*refereeAddress)

	skipIfis := parseSkipInterfaces()
	visionReceiver.MulticastServer.SkipInterfaces = skipIfis
	visionReceiver.MulticastServer.Verbose = *verbose
	trackedReceiver.MulticastServer.SkipInterfaces = skipIfis
	trackedReceiver.MulticastServer.Verbose = *verbose
	refereeReceiver.MulticastServer.SkipInterfaces = skipIfis
	refereeReceiver.MulticastServer.Verbose = *verbose

	visionReceiver.Start()
	trackedReceiver.Start()
	refereeReceiver.Start()

	srv := NewServer(
		visionReceiver.CombinedDetectionFrames,
		trackedReceiver.TrackedFrames,
		vision.GeometryProvider(visionReceiver),
		refereeReceiver.RefereeMsg,
	)
	return &http.Server{
		Addr:    *address,
		Handler: srv,
	}
}

func parseSkipInterfaces() []string {
	return strings.Split(*skipInterfaces, ",")
}

func formattedAddress() string {
	if strings.HasPrefix(*address, ":") {
		return "http://localhost" + *address
	}
	return "http://" + *address
}
