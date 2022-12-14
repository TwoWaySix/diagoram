package geometry

import (
	"math"
)

type BoundingBox struct {
	XMin, XMax, YMin, YMax float64
}

func NewBoundingBox(points []Point) BoundingBox {
	bb := BoundingBox{
		XMin: math.Inf(1),
		XMax: math.Inf(-1),
		YMin: math.Inf(1),
		YMax: math.Inf(-1),
	}
	for _, p := range points {
		if p.X < bb.XMin {
			bb.XMin = p.X
		}
		if p.X > bb.XMax {
			bb.XMax = p.X
		}
		if p.Y < bb.YMin {
			bb.YMin = p.Y
		}
		if p.Y > bb.YMax {
			bb.YMax = p.Y
		}
	}
	return bb
}
