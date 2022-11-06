package plotting

import "github.com/TwoWaySix/diagoram/geometry"

func RenderPoint(p geometry.Point, style MarkerStyle) []Pixel {
	var pixels []Pixel
	x0 := int(p.X) - style.Radius
	x1 := int(p.X) + style.Radius
	y0 := int(p.Y) - style.Radius
	y1 := int(p.Y) + style.Radius
	radiusPow := style.Radius * style.Radius

	for x := x0; x <= x1; x++ {
		for y := y0; y <= y1; y++ {
			dx := int(p.X) - x
			dy := int(p.Y) - y
			rPow := dx*dx + dy*dy
			if rPow < radiusPow {
				pixels = append(pixels, Pixel{
					X:     x,
					Y:     y,
					Color: style.Color,
				})
			}
		}
	}
	return pixels
}
