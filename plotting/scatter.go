package plotting

import (
	"github.com/TwoWaySix/diagoram/geometry"
)

type RenderingSettings struct {
	width, height uint
}

type ScatterPlot struct {
	points   []geometry.Point
	bBox     geometry.BoundingBox
	settings RenderingSettings
	style    MarkerStyle
}

func NewScatterPlot(series []geometry.Point, style MarkerStyle) ScatterPlot {
	sp := ScatterPlot{points: series}
	sp.bBox = geometry.NewBoundingBox(series)
	sp.style = style
	return sp
}
