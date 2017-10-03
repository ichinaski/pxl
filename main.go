package pxl

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <filename>...\n\n", os.Args[0])
		fmt.Println("Close the image with <ESC> or by pressing 'q'.")
		os.Exit(1)
	}

	Init()
	defer Close()
	for i := 1; i < len(os.Args); i++ {
		DisplayFile(os.Args[i])
	}
}
