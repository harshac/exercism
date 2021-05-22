package darts

import "math"

const (
	innerCircleRadius = 1
	innerCirclePoints = 10
	middleCircleRadius = 5
	middleCirclePoints = 5
	outerCircleRadius = 10
	outerCirclePoints = 1
	missedPoints = 0
)
func Score(x float64, y float64) int {
	dc := math.Hypot(x, y)
	if dc <= innerCircleRadius {
		return innerCirclePoints
	}
	if dc <= middleCircleRadius {
		return middleCirclePoints
	}
	if dc <= outerCircleRadius {
		return outerCirclePoints
	}
	return missedPoints
}
