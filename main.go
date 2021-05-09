// Package pxl is a little hack to display images in the terminal.
package pxl

import (
	"image"

	_ "image/jpeg"
	_ "image/png"

	"github.com/nsf/termbox-go"
)

// Draw clears terminal and draws the given image.
//
// Note: termbox.Init must have already been called.
func Draw(img image.Image) error {
	// Get terminal size and cursor width/height ratio
	width, height, whratio, err := canvasSize()
	if err != nil {
		return err
	}

	bounds := img.Bounds()
	imgW, imgH := bounds.Dx(), bounds.Dy()

	imgScale := scale(imgW, imgH, width, height, whratio)

	// Resize canvas to fit scaled image
	width, height = int(float64(imgW)/imgScale), int(float64(imgH)/(imgScale*whratio))

	err = termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		return err
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Calculate average color for the corresponding image rectangle
			// fitting in this cell. We use a half-block trick, wherein the
			// lower half of the cell displays the character ▄, effectively
			// doubling the resolution of the canvas.
			startX, startY, endX, endY := imgArea(x, y, imgScale, whratio)

			r, g, b := avgRGB(img, startX, startY, endX, (startY+endY)/2)
			colorUp := termbox.Attribute(termColor(r, g, b))

			r, g, b = avgRGB(img, startX, (startY+endY)/2, endX, endY)
			colorDown := termbox.Attribute(termColor(r, g, b))

			termbox.SetCell(x, y, '▄', colorDown, colorUp)
		}
	}
	return termbox.Flush()
}
