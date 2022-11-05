package plotting

import (
	"errors"
	"fmt"
	"github.com/TwoWaySix/diagoram/geometry"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"strings"
)

type RenderingSettings struct {
	width, height uint
}

type ScatterPlot struct {
	points   []geometry.Point
	bBox     geometry.BoundingBox
	settings RenderingSettings
}

func NewScatterPlot(series []geometry.Point) ScatterPlot {
	sp := ScatterPlot{points: series}
	sp.bBox = geometry.NewBoundingBox(series)
	return sp
}

func (sp *ScatterPlot) RenderToFile(fPath string, width, height int) error {
	img := createBlankImage(color.White, width, height)

	scaledPoints := sp.scalePoints(float64(width), float64(height))
	for _, p := range scaledPoints {
		pixels := RenderPoint(p)
		for _, pix := range pixels {
			img.Set(pix.X, pix.Y, p.Color)
		}
	}

	return writePixelsToFile(img, fPath)
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

func (sp *ScatterPlot) scalePoints(width, height float64) []geometry.Point {
	var scaledPoints []geometry.Point
	for _, p := range sp.points {
		relX := (p.X - sp.bBox.XMin) / (sp.bBox.XMax - sp.bBox.XMin)
		relY := (p.Y - sp.bBox.YMin) / (sp.bBox.YMax - sp.bBox.YMin)
		scaledPoints = append(scaledPoints, geometry.Point{
			X:      relX * width,
			Y:      relY * height,
			Color:  p.Color,
			Radius: p.Radius,
		})
	}
	return scaledPoints
}
