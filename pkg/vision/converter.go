package vision

import "github.com/RoboCup-SSL/ssl-go-tools/sslproto"

var white = "white"
var orange = "orange"
var yellow = "yellow"
var blue = "blue"
var lineWidth = 10
var ballRadius = float32(21)
var botRadius = float32(90)
var noFill = 0

func ProtoToPackage(frame *sslproto.SSL_DetectionFrame, geometry *sslproto.SSL_GeometryData) *Package {

	pack := new(Package)
	if geometry != nil {
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
	}

	for _, ball := range frame.Balls {
		pack.Circles = append(pack.Circles, Circle{
			Center: Point{*ball.X, *ball.Y},
			Radius: ballRadius,
			Style: Style{
				Fill: &orange,
			},
		})
	}

	for _, bot := range frame.RobotsBlue {
		pack.Circles = append(pack.Circles, Circle{
			Center: Point{*bot.X, *bot.Y},
			Radius: botRadius,
			Style: Style{
				Fill: &blue,
			},
		})
	}

	for _, bot := range frame.RobotsYellow {
		pack.Circles = append(pack.Circles, Circle{
			Center: Point{*bot.X, *bot.Y},
			Radius: botRadius,
			Style: Style{
				Fill: &yellow,
			},
		})
	}

	// TODO goals

	return pack
}
