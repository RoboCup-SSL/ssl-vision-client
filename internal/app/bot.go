package app

import "math"

type Bot struct {
	center2Dribbler float64
	botRadius       float64
}

func (b Bot) orient2CornerAngle() float64 {
	return math.Acos(b.center2Dribbler / b.botRadius)
}

func (b Bot) botRightX(orientation float64) float64 {
	return math.Cos(-orientation+b.orient2CornerAngle()) * b.botRadius
}

func (b Bot) botRightY(orientation float64) float64 {
	return math.Sin(-orientation+b.orient2CornerAngle()) * b.botRadius
}

func (b Bot) botLeftX(orientation float64) float64 {
	return math.Cos(-orientation-b.orient2CornerAngle()) * b.botRadius
}

func (b Bot) botLeftY(orientation float64) float64 {
	return math.Sin(-orientation-b.orient2CornerAngle()) * b.botRadius
}
