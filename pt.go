package main

import (
	"math"
)

type Ptf struct {
	X float64
	Y float64
}

func (p Ptf) DistanceTo(pb Ptf) float64 {
	dX := p.X - pb.X
	dY := p.Y - pb.Y
	return math.Sqrt(dX * dX + dY * dY)
}