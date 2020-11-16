package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/RoboCup-SSL/ssl-vision-client/pkg/vision"
	"github.com/golang/protobuf/proto"
	"log"
	"time"
)

var visionAddress = flag.String("visionAddress", "224.5.23.2:10006", "The multicast address of ssl-vision, default: 224.5.23.2:10006")
var fullScreen = flag.Bool("fullScreen", false, "Print the formatted message to the console, clearing the screen during print")

func main() {
	flag.Parse()

	receiver := vision.NewReceiver()
	go receiver.Receive(*visionAddress)

	for {
		if *fullScreen {
			// clear screen, move cursor to upper left corner
			fmt.Print("\033[H\033[2J")

			for camId, frame := range receiver.Detections() {
				fmt.Println("Camera ", camId)
				fmt.Println(proto.MarshalTextString(&frame))
				fmt.Println()
				fmt.Println()
			}
		} else {
			for _, frame := range receiver.Detections() {
				b, err := json.Marshal(frame)
				if err != nil {
					log.Fatal(err)
				}
				log.Print(string(b))
			}
		}

		time.Sleep(250 * time.Millisecond)
	}
}
