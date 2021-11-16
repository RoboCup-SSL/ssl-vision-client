package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/RoboCup-SSL/ssl-vision-client/pkg/vision"
	"github.com/golang/protobuf/proto"
	"log"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var visionAddress = flag.String("visionAddress", "224.5.23.2:10006", "The multicast address of ssl-vision, default: 224.5.23.2:10006")
var fullScreen = flag.Bool("fullScreen", false, "Print the formatted message to the console, clearing the screen during print")
var continuous = flag.Bool("continuous", false, "Print all received data in a continuous stream")
var noDetections = flag.Bool("noDetections", false, "Print the detection messages")
var noGeometry = flag.Bool("noGeometry", false, "Print the geometry messages")

func main() {
	flag.Parse()

	receiver := vision.NewReceiver()
	receiver.Start(*visionAddress)

	if *continuous {
		printContinuous(receiver)
	} else if *fullScreen {
		printFullscreen(receiver)
	} else {
		printAll(receiver)
	}
}

func printAll(receiver *vision.Receiver) {
	for {
		geometry := receiver.Geometry
		if !*noDetections {
			for _, frame := range receiver.Detections() {
				b, err := json.Marshal(frame)
				if err != nil {
					log.Fatal(err)
				}
				log.Print(string(b))
			}
		}
		if !*noGeometry && geometry != nil {
			b, err := json.Marshal(geometry)
			if err != nil {
				log.Fatal(err)
			}
			log.Print(string(b))
		}
		time.Sleep(250 * time.Millisecond)
	}
}

func printFullscreen(receiver *vision.Receiver) {
	for {
		geometry := receiver.Geometry
		// clear screen, move cursor to upper left corner
		fmt.Print("\033[H\033[2J")

		if !*noDetections {
			for camId, frame := range receiver.Detections() {
				fmt.Println("Camera ", camId)
				fmt.Println(proto.MarshalTextString(frame))
				fmt.Println()
				fmt.Println()
			}
		}
		if !*noGeometry && geometry != nil {
			fmt.Println(proto.MarshalTextString(geometry))
		}
		time.Sleep(250 * time.Millisecond)
	}
}

func printContinuous(receiver *vision.Receiver) {
	if !*noDetections {
		receiver.ConsumeDetections = func(frame *vision.SSL_DetectionFrame) {
			robots := append(frame.RobotsBlue, frame.RobotsYellow...)
			for _, robot := range robots {
				// ssl-vision may send a NaN confidence and the json serialization can not deal with it...
				if math.IsNaN(float64(*robot.Confidence)) {
					*robot.Confidence = 0
				}
			}
			b, err := json.Marshal(frame)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(b))
		}
	}
	if !*noGeometry {
		receiver.ConsumeGeometry = func(frame *vision.SSL_GeometryData) {
			b, err := json.Marshal(frame)
			if err != nil {
				log.Fatal(err)
			}
			log.Print(string(b))
		}
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
}
