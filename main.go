package main

import (
	"flag"
	"image"
	"os"
	"time"

	_ "image/jpeg"
	_ "image/png"

	"github.com/nsf/termbox-go"
)

var whratio float64 // The terminal's cursor width/height ratio

func init() {
	flag.Float64Var(&whratio, "r", 2.35, "Cursor width/height ratio in your terminal")
}

func draw(img image.Image) {
	w, h := termbox.Size()

	bounds := img.Bounds()
	imgW, imgH := bounds.Dx(), bounds.Dy()

	imgScale := scale(imgW, imgH, w, h, whratio)

	// Resize canvas to fit scaled image
	w, h = int(float64(imgW)/imgScale), int(float64(imgH)/(imgScale*whratio))

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
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
	termbox.Flush()
}

func main() {
	flag.Parse()
	img, err := load(os.Args[len(os.Args)-1])
	if err != nil {
		panic(err)
	}

	err = termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetOutputMode(termbox.Output256)

	draw(img)
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break loop
			}
		case termbox.EventResize:
			draw(img)
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}
