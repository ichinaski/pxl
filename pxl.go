package pxl

import (
	"fmt"
	"image"
	"time"

	_ "image/jpeg"
	_ "image/png"

	"github.com/nsf/termbox-go"
)

func Draw(img image.Image) {
	// Get terminal size and cursor width/height ratio
	width, height, whratio := canvasSize()

	bounds := img.Bounds()
	imgW, imgH := bounds.Dx(), bounds.Dy()

	imgScale := scale(imgW, imgH, width, height, whratio)

	// Resize canvas to fit scaled image
	width, height = int(float64(imgW)/imgScale), int(float64(imgH)/(imgScale*whratio))

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
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
	termbox.Flush()
}

func DisplayFile(image string) {
	img, err := load(image)
	if err != nil {
		panic(err)
	}
	Display(img)
}

func Display(img image.Image) {
	Draw(img)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc || ev.Ch == 'q' {
				return
			}
		case termbox.EventResize:
			Draw(img)
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func Init() {
	fmt.Println("FOO")
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetOutputMode(termbox.Output256)
}

func Close() {
	termbox.Close()
}
