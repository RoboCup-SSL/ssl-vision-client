package vision

func GeometryProvider(receiver *Receiver) func() *SSL_GeometryData {
	return func() *SSL_GeometryData {
		geometry := receiver.CurrentGeometry()
		if geometry == nil {
			return defaultGeometry()
		}
		return geometry
	}
}

func defaultGeometry() (g *SSL_GeometryData) {
	g = new(SSL_GeometryData)
	g.Field = new(SSL_GeometryFieldSize)
	g.Field.FieldLength = new(int32)
	g.Field.FieldWidth = new(int32)
	g.Field.GoalWidth = new(int32)
	g.Field.GoalDepth = new(int32)
	g.Field.BoundaryWidth = new(int32)
	g.Field.PenaltyAreaDepth = new(int32)
	g.Field.PenaltyAreaWidth = new(int32)
	g.Field.CenterCircleRadius = new(int32)
	g.Field.LineThickness = new(int32)
	g.Field.GoalCenterToPenaltyMark = new(int32)
	g.Field.GoalHeight = new(int32)
	g.Field.BallRadius = new(float32)
	g.Field.MaxRobotRadius = new(float32)

	*g.Field.FieldLength = 12000
	*g.Field.FieldWidth = 9000
	*g.Field.GoalWidth = 1000
	*g.Field.GoalDepth = 180
	*g.Field.BoundaryWidth = 300
	*g.Field.PenaltyAreaDepth = 1800
	*g.Field.PenaltyAreaWidth = 3600
	*g.Field.CenterCircleRadius = 500
	*g.Field.LineThickness = 10
	*g.Field.GoalCenterToPenaltyMark = 8000
	*g.Field.GoalHeight = 155
	*g.Field.BallRadius = 10
	*g.Field.MaxRobotRadius = 90
	return
}
