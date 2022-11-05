package plotting

import "github.com/TwoWaySix/diagoram/geometry"

type ScatterPlot struct {
	series      []geometry.Point2D
	boundingBox geometry.BoundingBox
}

func NewScatterPlot(series []geometry.Point2D) ScatterPlot {
	sp := ScatterPlot{series: series}
	sp.boundingBox = geometry.NewBoundingBox(series)
	return sp
}
