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
