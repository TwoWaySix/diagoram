package plotting

import "github.com/TwoWaySix/diagoram/geometry"

func RenderPoint(p geometry.Point) []Pixel {
	var pixels []Pixel
	x0 := int(p.X - p.Radius)
	x1 := int(p.X + p.Radius)
	y0 := int(p.Y - p.Radius)
	y1 := int(p.Y + p.Radius)
	radiusPow := p.Radius * p.Radius

	for x := x0; x <= x1; x++ {
		for y := y0; y <= y1; y++ {
			dx := p.X - float64(x)
			dy := p.Y - float64(y)
			rPow := dx*dx + dy*dy
			if rPow < radiusPow {
				pixels = append(pixels, Pixel{
					X:     x,
					Y:     y,
					Color: p.Color,
				})
			}
		}
	}
	return pixels
}
