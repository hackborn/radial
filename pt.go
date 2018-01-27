package main

import (
	"math"
)

// Ptf represents a floating 2D point.
type Ptf struct {
	X float64
	Y float64
}

// DistanceTo provides the distance between points.
func (p Ptf) DistanceTo(pb Ptf) float64 {
	dX := p.X - pb.X
	dY := p.Y - pb.Y
	return math.Sqrt(dX*dX + dY*dY)
}
