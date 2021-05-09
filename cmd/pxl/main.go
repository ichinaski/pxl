// Command pxl is a little hack to display images in the terminal.
package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"os"
	"time"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	// If we wanted to include golang.org/x/image as a
	// dependency we could also support:
	//_ "golang.org/x/image/bmp"
	//_ "golang.org/x/image/tiff"
	//_ "golang.org/x/image/webp"

	"github.com/ichinaski/pxl"
	"github.com/nsf/termbox-go"
)

// display a single image.
// It redraws on resize events and waits for
// a key event of either ESC or `q` before returning.
func display(image string) error {
	img, err := load(image)
	if err != nil {
		return err
	}
	if err := pxl.Draw(img); err != nil {
		return err
	}

	// Wait for keypress or resize.
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc || ev.Ch == 'q' {
				return nil
			}
		case termbox.EventResize:
			err = pxl.Draw(img)
			if err != nil {
				return err
			}
		default:
			// PollEvent blocks, but we'll do a small sleep for safety.
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func main() {
	log.SetPrefix(os.Args[0] + ": ")
	log.SetFlags(0)
	flag.Usage = func() {
		out := flag.CommandLine.Output()
		fmt.Fprintf(out, "Usage: %s <filename>...\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(out, "\nClose the image with <ESC> or by pressing 'q'.")
		os.Exit(66) // EX_USAGE
	}
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
	}

	err := termbox.Init()
	if err != nil {
		log.Fatalln("Failed to initialise terminal:", err)
	}
	// safe to call repeatedly; defer to guard against panics
	defer termbox.Close()

	errs := displayAll(flag.Args()...)
	termbox.Close()

	// We've closed termbox so now we're safe to use log again
	// in order to report any errors that may have been encountered.
	var good, bad int
	for _, err := range errs {
		if err != nil {
			bad++
			log.Println(err)
		} else {
			good++
		}
	}
	if bad > 0 {
		if good > 0 {
			os.Exit(1) // Some worked, 1 is "alternative success"
		}
		os.Exit(66) // EX_NOINPUT, anything >=2 is a failure
	}
}

// displayAll calls display for each path.
// It saves any errors until after the last image so that they
// can be shown after the termbox is closed.
func displayAll(paths ...string) []error {
	termbox.SetOutputMode(termbox.Output256)

	errs := make([]error, 0, len(paths))
	for _, arg := range flag.Args() {
		err := display(arg)
		errs = append(errs, err)
	}
	return errs
}

// load an image stored in the given path
func load(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	return img, err
}
