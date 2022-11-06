package plotting

import (
	"errors"
	"fmt"
	"github.com/TwoWaySix/diagoram/geometry"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
	"strings"
)

type Plot struct {
	title       string
	series      []ScatterPlot
	bBox        geometry.BoundingBox
	graphMargin Margin
}

type Margin struct {
	Top, Right, Bottom, Left float64
}

func NewPlot(title string) Plot {
	plt := Plot{title: title}
	plt.graphMargin = Margin{
		Top:    0.15,
		Right:  0.05,
		Bottom: 0.1,
		Left:   0.05,
	}
	return plt
}

func (plt *Plot) AddSeries(sp ScatterPlot) {
	plt.series = append(plt.series, sp)
}

func (plt *Plot) RenderToFile(fPath string, width int, height int) error {
	plt.updateBoundingBox()

	img := createBlankImage(color.White, width, height)
	plt.renderAxes(float64(width), float64(height), img)
	plt.renderSeries(width, height, img)
	return writePixelsToFile(img, fPath)
}

func (plt *Plot) renderSeries(width int, height int, img *image.RGBA) {
	for _, sp := range plt.series {
		scaledPoints := scalePoints(sp.points, plt.bBox, float64(width), float64(height), plt.graphMargin)
		for _, p := range scaledPoints {
			pixels := RenderPoint(p, sp.style)
			for _, pix := range pixels {
				img.Set(pix.X, pix.Y, sp.style.Color)
			}
		}
	}
}

func (plt *Plot) updateBoundingBox() {
	bb := geometry.BoundingBox{
		XMin: math.Inf(1),
		XMax: math.Inf(-1),
		YMin: math.Inf(1),
		YMax: math.Inf(-1),
	}
	for _, sp := range plt.series {
		for _, p := range sp.points {
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
	}
	plt.bBox = bb
}

func (plt *Plot) renderAxes(width float64, height float64, img *image.RGBA) {
	x0 := int(width * plt.graphMargin.Left)
	x1 := int(width - width*plt.graphMargin.Right)
	y1 := int(height * plt.graphMargin.Top)
	y0 := int(height - height*plt.graphMargin.Bottom)

	for x := x0; x < x1; x++ {
		img.Set(x, y0, color.Black)
	}
	for y := y1; y < y0; y++ {
		img.Set(x0, y, color.Black)
	}
}

func scalePoints(points []geometry.Point, bBox geometry.BoundingBox, width float64, height float64, margin Margin) []geometry.Point {
	x0 := width * margin.Left
	x1 := width - width*margin.Right
	y1 := height * margin.Top
	y0 := height - height*margin.Bottom

	var scaledPoints []geometry.Point
	for _, p := range points {
		relX := (p.X - bBox.XMin) / (bBox.XMax - bBox.XMin)
		relY := (p.Y - bBox.YMin) / (bBox.YMax - bBox.YMin)

		scaledPoints = append(scaledPoints, geometry.Point{
			X: x0 + relX*(x1-x0),
			Y: y1 + relY*(y0-y1),
		})
	}
	return scaledPoints
}

func createBlankImage(bgColor color.Color, width, height int) *image.RGBA {
	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{},
		Max: image.Point{X: width, Y: height},
	})
	draw.Draw(img, img.Bounds(), &image.Uniform{C: bgColor}, image.Point{}, draw.Src)
	return img
}

func writePixelsToFile(img *image.RGBA, fPath string) error {
	if strings.HasSuffix(strings.ToUpper(fPath), "PNG") {
		f, _ := os.Create(fPath)
		err := png.Encode(f, img)
		if err != nil {
			return fmt.Errorf("rendering to file: %v", err)
		}
	} else {
		return errors.New("wrong output format: allowed formats: PNG")
	}
	return nil
}
