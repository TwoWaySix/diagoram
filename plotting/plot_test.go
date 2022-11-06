package plotting

import (
	"github.com/TwoWaySix/diagoram/geometry"
	"github.com/stretchr/testify/assert"
	"image/color"
	"testing"
)

func TestPlot_UpdatingBoundingBox(t *testing.T) {
	// Arrange
	sp1 := NewScatterPlot([]geometry.Point{
		{
			X: 10,
			Y: 11,
		}, {
			X: 5,
			Y: 6,
		},
	}, MarkerStyle{
		Radius: 5,
		Color:  color.Black,
	})
	sp2 := NewScatterPlot([]geometry.Point{
		{
			X: 0,
			Y: 1,
		}, {
			X: 5,
			Y: 6,
		},
	}, MarkerStyle{
		Radius: 5,
		Color:  color.RGBA{R: 255, A: 255},
	})

	plt := NewPlot("dummy title")
	plt.AddSeries(sp1)
	plt.AddSeries(sp2)

	// Act
	plt.updateBoundingBox()

	// Assert
	assert.Equal(t, plt.bBox.XMin, 0.)
	assert.Equal(t, plt.bBox.XMax, 10.)
	assert.Equal(t, plt.bBox.YMin, 1.)
	assert.Equal(t, plt.bBox.YMax, 11.)
}

func TestPlot_RenderToFile(t *testing.T) {
	// Arrange
	plt := NewPlot("dummy title")
	style := MarkerStyle{
		Radius: 5,
		Color:  color.RGBA{R: 255, A: 255},
	}
	sp := NewScatterPlot([]geometry.Point{
		{
			X: 0,
			Y: 1,
		}, {
			X: 10,
			Y: 11,
		}, {
			X: 5,
			Y: 6,
		}, {
			X: 0,
			Y: 11,
		}, {
			X: 10,
			Y: 1,
		},
	}, style)
	plt.series = append(plt.series, sp)

	// Act
	err := plt.RenderToFile("./testdata/corner_points_and_center.png", 800, 600)

	// Assert
	assert.Nil(t, err)
}
