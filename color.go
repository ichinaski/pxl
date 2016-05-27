package main

import (
	"image/color"
	"math"
)

func rgb(c color.Color) (uint16, uint16, uint16) {
	r, g, b, _ := c.RGBA()
	// Reduce color values to the range [0, 15]
	return uint16(r >> 8), uint16(g >> 8), uint16(b >> 8)
}

// termColor converts a 24-bit RGB color into a term256 compatible approximation.
func termColor(r, g, b float64) uint16 {
	rterm := uint16(math.Floor(r*5+0.5)) * 36
	gterm := uint16(math.Floor(g*5+0.5)) * 6
	bterm := uint16(math.Floor(b*5 + 0.5))

	return rterm + gterm + bterm + 16 + 1 // termbox default color offset
}
