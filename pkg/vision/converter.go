package vision

import (
	"github.com/RoboCup-SSL/ssl-go-tools/pkg/sslproto"
	"math"
	"strconv"
)

var white = "white"
var orange = "orange"
var yellow = "yellow"
var blue = "blue"
var black = "black"
var lineWidth = float32(10)
var ballRadius = float32(21)
var botRadius = float64(90)
var center2Dribbler = float64(75)
var noFill = float32(0)
var botStrokeWidth = float32(10)
var ballStrokeWidth = float32(0)

func ProtoToPackage(frame *sslproto.SSL_DetectionFrame, geometry *sslproto.SSL_GeometryData) *Package {

	pack := new(Package)
	if geometry != nil {
		addGeometryShapes(pack, geometry)
	}

	for _, ball := range frame.Balls {
		pack.Circles = append(pack.Circles, createBallShape(ball))
	}

	for _, bot := range frame.RobotsBlue {
		pack.Paths = append(pack.Paths, createBotPath(bot, blue))
		pack.Texts = append(pack.Texts, createBotId(bot, white))
	}

	for _, bot := range frame.RobotsYellow {
		pack.Paths = append(pack.Paths, createBotPath(bot, yellow))
		pack.Texts = append(pack.Texts, createBotId(bot, black))
	}

	return pack
}

func createBallShape(ball *sslproto.SSL_DetectionBall) Circle {
	return Circle{
		Center: Point{*ball.X, *ball.Y},
		Radius: ballRadius,
		Style: Style{
			StrokeWidth: &ballStrokeWidth,
			Fill:        &orange,
		},
	}
}

func addGeometryShapes(pack *Package, geometry *sslproto.SSL_GeometryData) {
	pack.FieldWidth = float32(*geometry.Field.FieldWidth)
	pack.FieldLength = float32(*geometry.Field.FieldLength)
	pack.BoundaryWidth = float32(*geometry.Field.BoundaryWidth)
	pack.GoalWidth = float32(*geometry.Field.GoalWidth)
	pack.GoalDepth = float32(*geometry.Field.GoalDepth)
	for _, line := range geometry.Field.FieldLines {
		pack.Lines = append(pack.Lines, Line{
			P1: Point{*line.P1.X, *line.P1.Y},
			P2: Point{*line.P2.X, *line.P2.Y},
			Style: Style{
				Stroke:      &white,
				StrokeWidth: &lineWidth,
			},
		})
	}
	for _, arc := range geometry.Field.FieldArcs {
		pack.Circles = append(pack.Circles, Circle{
			Center: Point{*arc.Center.X, *arc.Center.Y},
			Radius: *arc.Radius,
			Style: Style{
				Stroke:      &white,
				StrokeWidth: &lineWidth,
				FillOpacity: &noFill,
			},
		})
	}
	pack.Lines = append(pack.Lines, goalLinesPositive(geometry)...)
	pack.Lines = append(pack.Lines, goalLinesNegative(geometry)...)
}

func goalLinesNegative(geometry *sslproto.SSL_GeometryData) (lines []Line) {
	lines = goalLinesPositive(geometry)
	for i := range lines {
		lines[i].P1.X *= -1
		lines[i].P2.X *= -1
	}
	return
}

func goalLinesPositive(geometry *sslproto.SSL_GeometryData) (lines []Line) {
	flh := float32(*geometry.Field.FieldLength / 2)
	gwh := float32(*geometry.Field.GoalWidth / 2)
	gd := float32(*geometry.Field.GoalDepth)

	lines = append(lines, Line{P1: Point{flh, -gwh}, P2: Point{flh + gd, -gwh},
		Style: Style{Stroke: &black, StrokeWidth: &lineWidth}})
	lines = append(lines, Line{P1: Point{flh, gwh}, P2: Point{flh + gd, gwh},
		Style: Style{Stroke: &black, StrokeWidth: &lineWidth}})
	lines = append(lines, Line{P1: Point{flh + gd, -gwh}, P2: Point{flh + gd, gwh},
		Style: Style{Stroke: &black, StrokeWidth: &lineWidth}})
	return
}

func createBotPath(bot *sslproto.SSL_DetectionRobot, fillColor string) Path {
	b := Bot{center2Dribbler, botRadius}
	x := float64(*bot.X)
	y := float64(*bot.Y)
	o := float64(*bot.Orientation)
	return Path{
		D: []PathElement{
			{Type: "M",
				Args: []float64{
					x + b.botRightX(o),
					y + b.botRightY(o),
				},
			},
			{Type: "A",
				Args: []float64{
					botRadius,
					botRadius,
					0, 1, 1,
					x + b.botLeftX(o),
					y + b.botLeftY(o),
				},
			},
			{Type: "L",
				Args: []float64{
					x + b.botRightX(o),
					y + b.botRightY(o),
				},
			},
		},
		Style: Style{
			Fill:        &fillColor,
			Stroke:      &black,
			StrokeWidth: &botStrokeWidth,
		},
	}
}

func createBotId(bot *sslproto.SSL_DetectionRobot, strokeColor string) Text {
	return Text{
		Text: strconv.Itoa(int(*bot.RobotId)),
		P:    Point{*bot.X, *bot.Y},
		Style: Style{
			Fill: &strokeColor,
		},
	}
}

type Bot struct {
	center2Dribbler float64
	botRadius       float64
}

func (b Bot) orient2CornerAngle() float64 {
	return math.Acos(b.center2Dribbler / b.botRadius)
}

func (b Bot) botRightX(orientation float64) float64 {
	return math.Cos(orientation+b.orient2CornerAngle()) * b.botRadius
}

func (b Bot) botRightY(orientation float64) float64 {
	return math.Sin(orientation+b.orient2CornerAngle()) * b.botRadius
}

func (b Bot) botLeftX(orientation float64) float64 {
	return math.Cos(orientation-b.orient2CornerAngle()) * b.botRadius
}

func (b Bot) botLeftY(orientation float64) float64 {
	return math.Sin(orientation-b.orient2CornerAngle()) * b.botRadius
}
