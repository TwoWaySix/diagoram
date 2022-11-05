package plotting

import (
	"github.com/TwoWaySix/diagoram/geometry"
	"github.com/stretchr/testify/assert"
	"image/color"
	"testing"
)

func TestNewScatterPlot(t *testing.T) {
	sp := NewScatterPlot([]geometry.Point{
		{
			X:      0,
			Y:      1,
			Color:  color.RGBA{R: 255, A: 255},
			Radius: 10,
		}, {
			X:      10,
			Y:      11,
			Color:  color.RGBA{B: 255, A: 255},
			Radius: 10,
		}, {
			X:      5,
			Y:      6,
			Color:  color.RGBA{G: 255, A: 255},
			Radius: 10,
		},
	})
	assert.Equal(t, sp.bBox.XMin, 0.)
	assert.Equal(t, sp.bBox.XMax, 10.)
	assert.Equal(t, sp.bBox.YMin, 1.)
	assert.Equal(t, sp.bBox.YMax, 11.)
}

func TestScatterPlot_RenderToFile(t *testing.T) {
	sp := NewScatterPlot([]geometry.Point{
		{
			X:      0,
			Y:      1,
			Color:  color.RGBA{R: 255, A: 255},
			Radius: 10,
		}, {
			X:      10,
			Y:      11,
			Color:  color.RGBA{B: 255, A: 255},
			Radius: 10,
		}, {
			X:      5,
			Y:      6,
			Color:  color.RGBA{G: 255, A: 255},
			Radius: 10,
		},
	})
	err := sp.RenderToFile("./testdata/three_points.png", 400, 200)
	assert.Nil(t, err)
}

func TestScalePoints(t *testing.T) {

}
