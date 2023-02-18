package client

import (
	"fmt"
	"github.com/RoboCup-SSL/ssl-vision-client/pkg/tracked"
	"github.com/RoboCup-SSL/ssl-vision-client/pkg/vision"
	"github.com/RoboCup-SSL/ssl-vision-client/pkg/visualization"
	"math"
	"sort"
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

type ShapesByOrderNumber []Shape

func (a ShapesByOrderNumber) Len() int           { return len(a) }
func (a ShapesByOrderNumber) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ShapesByOrderNumber) Less(i, j int) bool { return a[i].OrderNumber < a[j].OrderNumber }

func (p *Package) AddDetectionFrame(frame *vision.SSL_DetectionFrame) {
	for _, ball := range frame.Balls {
		p.Shapes = append(p.Shapes, Shape{OrderNumber: 3, Circle: createBallShape(*ball.X, *ball.Y, 0)})
	}

	for _, bot := range frame.RobotsBlue {
		p.Shapes = append(p.Shapes, Shape{OrderNumber: 1, Path: createBotPath(*bot.X, *bot.Y, *bot.Orientation, blue)})
		p.Shapes = append(p.Shapes, Shape{OrderNumber: 2, Text: createBotId(*bot.RobotId, *bot.X, *bot.Y, white)})
	}

	for _, bot := range frame.RobotsYellow {
		p.Shapes = append(p.Shapes, Shape{OrderNumber: 1, Path: createBotPath(*bot.X, *bot.Y, *bot.Orientation, yellow)})
		p.Shapes = append(p.Shapes, Shape{OrderNumber: 2, Text: createBotId(*bot.RobotId, *bot.X, *bot.Y, black)})
	}
}

func (p *Package) AddTrackedFrame(frame *tracked.TrackerWrapperPacket) {
	if frame.TrackedFrame == nil {
		return
	}
	for _, ball := range frame.TrackedFrame.Balls {
		p.Shapes = append(p.Shapes, Shape{OrderNumber: 3, Circle: createBallShape(*ball.Pos.X*1000, *ball.Pos.Y*1000, *ball.Pos.Z*1000)})
		p.Shapes = append(p.Shapes, Shape{OrderNumber: 1, Circle: createBallHighlighter(*ball.Pos.X*1000, *ball.Pos.Y*1000)})
	}

	for _, bot := range frame.TrackedFrame.Robots {
		var botColor string
		var strokeColor string
		if *bot.RobotId.Team == tracked.Team_YELLOW {
			botColor = yellow
			strokeColor = black
		} else {
			botColor = blue
			strokeColor = white
		}
		p.Shapes = append(p.Shapes, Shape{OrderNumber: 1, Path: createBotPath(*bot.Pos.X*1000, *bot.Pos.Y*1000, *bot.Orientation, botColor)})
		p.Shapes = append(p.Shapes, Shape{OrderNumber: 2, Text: createBotId(*bot.RobotId.Id, *bot.Pos.X*1000, *bot.Pos.Y*1000, strokeColor)})
	}
}

func createBallShape(x, y, z float32) *Circle {
	heightFactor := 0.01*math.Abs(float64(z)) + 1
	return &Circle{
		Center: Point{x, -y},
		Radius: float32(heightFactor) * ballRadius,
		Style: Style{
			StrokeWidth: &ballStrokeWidth,
			Fill:        &orange,
		},
	}
}

func createBallHighlighter(x, y float32) *Circle {
	return &Circle{
		Center: Point{x, -y},
		Radius: 500,
		Style: Style{
			Stroke:      &orange,
			StrokeWidth: &lineWidth,
			FillOpacity: &noFill,
		},
	}
}

func (p *Package) AddGeometryShapes(geometry *vision.SSL_GeometryData) {
	p.FieldWidth = float32(*geometry.Field.FieldWidth)
	p.FieldLength = float32(*geometry.Field.FieldLength)
	p.BoundaryWidth = float32(*geometry.Field.BoundaryWidth)
	p.GoalWidth = float32(*geometry.Field.GoalWidth)
	p.GoalDepth = float32(*geometry.Field.GoalDepth)
	for _, line := range geometry.Field.FieldLines {
		p.Shapes = append(p.Shapes, Shape{
			Line: &Line{
				P1: Point{*line.P1.X, -*line.P1.Y},
				P2: Point{*line.P2.X, -*line.P2.Y},
				Style: Style{
					Stroke:      &white,
					StrokeWidth: &lineWidth,
				},
			},
		})
	}
	for _, arc := range geometry.Field.FieldArcs {
		p.Shapes = append(p.Shapes, Shape{
			Circle: &Circle{
				Center: Point{*arc.Center.X, -*arc.Center.Y},
				Radius: *arc.Radius,
				Style: Style{
					Stroke:      &white,
					StrokeWidth: &lineWidth,
					FillOpacity: &noFill,
				},
			},
		})
	}
	p.Shapes = append(p.Shapes, goalLinesPositive(geometry)...)
	p.Shapes = append(p.Shapes, goalLinesNegative(geometry)...)
}

func goalLinesNegative(geometry *vision.SSL_GeometryData) (lines []Shape) {
	lines = goalLinesPositive(geometry)
	for i := range lines {
		lines[i].Line.P1.X *= -1
		lines[i].Line.P2.X *= -1
	}
	return
}

func goalLinesPositive(geometry *vision.SSL_GeometryData) (lines []Shape) {
	flh := float32(*geometry.Field.FieldLength / 2)
	gwh := float32(*geometry.Field.GoalWidth / 2)
	gd := float32(*geometry.Field.GoalDepth)

	lines = append(lines, Shape{Line: &Line{P1: Point{flh, -gwh}, P2: Point{flh + gd, -gwh},
		Style: Style{Stroke: &black, StrokeWidth: &lineWidth}}})
	lines = append(lines, Shape{Line: &Line{P1: Point{flh, gwh}, P2: Point{flh + gd, gwh},
		Style: Style{Stroke: &black, StrokeWidth: &lineWidth}}})
	lines = append(lines, Shape{Line: &Line{P1: Point{flh + gd, -gwh}, P2: Point{flh + gd, gwh},
		Style: Style{Stroke: &black, StrokeWidth: &lineWidth}}})
	return
}

func createBotPath(posX, posY, orientation float32, fillColor string) *Path {
	b := Bot{center2Dribbler, botRadius}
	x := float64(posX)
	y := -float64(posY)
	o := float64(orientation)
	return &Path{
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

func createBotId(id uint32, x, y float32, strokeColor string) *Text {
	strokeWidth := float32(0)
	return &Text{
		Text: strconv.Itoa(int(id)),
		P:    Point{x, -y},
		Style: Style{
			Stroke:      &strokeColor,
			Fill:        &strokeColor,
			StrokeWidth: &strokeWidth,
		},
	}
}

func (p *Package) SortShapes() {
	sort.Sort(ShapesByOrderNumber(p.Shapes))
}

func (p *Package) AddLineSegment(sourceId string, lineSegment *visualization.LineSegment) {
	p.Shapes = append(p.Shapes, Shape{
		OrderNumber: lineSegment.Metadata.Order,
		Line: &Line{
			P1: Point{lineSegment.StartX * 1000, lineSegment.StartY * 1000},
			P2: Point{lineSegment.EndX * 1000, lineSegment.EndY * 1000},
			Metadata: Metadata{
				SourceId:         sourceId,
				Layer:            lineSegment.Metadata.Layer,
				VisibleByDefault: lineSegment.Metadata.VisibleByDefault,
			},
			Style: Style{
				Fill:   rgb(lineSegment.Metadata.ColorFill),
				Stroke: rgb(lineSegment.Metadata.ColorStroke),
			},
		},
	})
}

func (p *Package) AddCircle(sourceId string, circle *visualization.Circle) {
	p.Shapes = append(p.Shapes, Shape{
		OrderNumber: circle.Metadata.Order,
		Circle: &Circle{
			Center: Point{circle.CenterX * 1000, circle.CenterY * 1000},
			Radius: circle.Radius * 1000,
			Metadata: Metadata{
				SourceId:         sourceId,
				Layer:            circle.Metadata.Layer,
				VisibleByDefault: circle.Metadata.VisibleByDefault,
			},
			Style: Style{
				Fill:   rgb(circle.Metadata.ColorFill),
				Stroke: rgb(circle.Metadata.ColorStroke),
			},
		},
	})
}

func rgb(rgb *visualization.RgbColor) *string {
	if rgb == nil {
		return nil
	}
	color := fmt.Sprintf("rgba(%d,%d,%d,%.5f)", rgb.R, rgb.G, rgb.B, rgb.A)
	return &color
}
