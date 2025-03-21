package main

import (
	"flag"
	"github.com/RoboCup-SSL/ssl-vision-client/frontend"
	"github.com/RoboCup-SSL/ssl-vision-client/pkg/client"
	"github.com/RoboCup-SSL/ssl-vision-client/pkg/referee"
	"github.com/RoboCup-SSL/ssl-vision-client/pkg/tracked"
	"github.com/RoboCup-SSL/ssl-vision-client/pkg/vision"
	"github.com/RoboCup-SSL/ssl-vision-client/pkg/visualization"
	"log"
	"net/http"
	"strings"
)

var address = flag.String("address", ":8082", "The address on which the UI and API is served, default: :8082")
var visionAddress = flag.String("visionAddress", "224.5.23.2:10006", "The multicast address of ssl-vision, default: 224.5.23.2:10006")
var trackedAddress = flag.String("trackedAddress", "224.5.23.2:10010", "The multicast address of trackers, default: 224.5.23.2:10010")
var visualizationAddress = flag.String("visualizationAddress", "224.5.23.2:10012", "The multicast address of visualization frames, default: 224.5.23.2:10012")
var refereeAddress = flag.String("refereeAddress", "224.5.23.1:10003", "The multicast address of the game controller, default: 224.5.23.1:10003")
var skipInterfaces = flag.String("skipInterfaces", "", "Comma separated list of interface names to ignore when receiving multicast packets")
var verbose = flag.Bool("verbose", false, "Verbose output")

func main() {
	flag.Parse()

	setupVisionClient()
	frontend.HandleUi()

	log.Printf("UI is available at %v", formattedAddress())

	err := http.ListenAndServe(*address, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func setupVisionClient() {
	visionReceiver := vision.NewReceiver(*visionAddress)
	visualizationReceiver := visualization.NewReceiver(*visualizationAddress)
	trackedReceiver := tracked.NewReceiver(*trackedAddress)
	refereeReceiver := referee.NewReceiver(*refereeAddress)

	publisher := client.NewPublisher()
	publisher.DetectionProvider = visionReceiver.CombinedDetectionFrames
	publisher.TrackerProvider = trackedReceiver.TrackedFrames
	publisher.GeometryProvider = geometryProvider(visionReceiver)
	publisher.RefereeProvider = refereeReceiver.RefereeMsg
	publisher.LineSegmentProvider = visualizationReceiver.GetLineSegments
	publisher.CircleProvider = visualizationReceiver.GetCircles
	http.HandleFunc("/api/vision", publisher.Handler)

	skipIfis := parseSkipInterfaces()
	visionReceiver.MulticastServer.SkipInterfaces = skipIfis
	visionReceiver.MulticastServer.Verbose = *verbose
	visualizationReceiver.MulticastServer.SkipInterfaces = skipIfis
	visualizationReceiver.MulticastServer.Verbose = *verbose
	trackedReceiver.MulticastServer.SkipInterfaces = skipIfis
	trackedReceiver.MulticastServer.Verbose = *verbose
	refereeReceiver.MulticastServer.SkipInterfaces = skipIfis
	refereeReceiver.MulticastServer.Verbose = *verbose

	visionReceiver.Start()
	visualizationReceiver.Start()
	trackedReceiver.Start()
	refereeReceiver.Start()
}

func geometryProvider(receiver *vision.Receiver) func() *vision.SSL_GeometryData {
	return func() *vision.SSL_GeometryData {
		geometry := receiver.CurrentGeometry()
		if geometry == nil {
			return defaultGeometry()
		}
		return geometry
	}
}

func defaultGeometry() (g *vision.SSL_GeometryData) {
	g = new(vision.SSL_GeometryData)
	g.Field = new(vision.SSL_GeometryFieldSize)
	g.Field.FieldWidth = new(int32)
	g.Field.FieldLength = new(int32)
	g.Field.GoalDepth = new(int32)
	g.Field.GoalWidth = new(int32)
	g.Field.BoundaryWidth = new(int32)
	*g.Field.FieldWidth = 9000
	*g.Field.FieldLength = 12000
	*g.Field.GoalDepth = 180
	*g.Field.GoalWidth = 1000
	*g.Field.BoundaryWidth = 300
	return
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
