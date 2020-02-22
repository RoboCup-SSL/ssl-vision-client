package client

import (
	"fmt"
	"github.com/RoboCup-SSL/ssl-vision-client/pkg/vision"
	"github.com/RoboCup-SSL/ssl-vision-client/pkg/visualization"
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
		p.Shapes = append(p.Shapes, Shape{OrderNumber: 3, Circle: createBallShape(ball)})
	}

	for _, bot := range frame.RobotsBlue {
		p.Shapes = append(p.Shapes, Shape{OrderNumber: 1, Path: createBotPath(bot, blue)})
		p.Shapes = append(p.Shapes, Shape{OrderNumber: 2, Text: createBotId(bot, white)})
	}

	for _, bot := range frame.RobotsYellow {
		p.Shapes = append(p.Shapes, Shape{OrderNumber: 1, Path: createBotPath(bot, yellow)})
		p.Shapes = append(p.Shapes, Shape{OrderNumber: 2, Text: createBotId(bot, black)})
	}
}

func createBallShape(ball *vision.SSL_DetectionBall) *Circle {
	return &Circle{
		Center: Point{*ball.X, -*ball.Y},
		Radius: ballRadius,
		Style: Style{
			StrokeWidth: &ballStrokeWidth,
			Fill:        &orange,
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

func createBotPath(bot *vision.SSL_DetectionRobot, fillColor string) *Path {
	b := Bot{center2Dribbler, botRadius}
	x := float64(*bot.X)
	y := -float64(*bot.Y)
	o := float64(*bot.Orientation)
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

func createBotId(bot *vision.SSL_DetectionRobot, strokeColor string) *Text {
	return &Text{
		Text: strconv.Itoa(int(*bot.RobotId)),
		P:    Point{*bot.X, -*bot.Y},
		Style: Style{
			Fill: &strokeColor,
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
