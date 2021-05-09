// Builds for these work but not all the resultant binaries have been tested.
// +build darwin dragonfly freebsd linux netbsd openbsd

// Note that as of at least termbox-go@v1.1.1 the following GOOS
// do not build (for termbox):
//	aix
//	illumos
//	plan9
//	solaris
// For other reasons, pxl doesn't appear to build on GOOS:
//	android
//	ios

package main

import (
	"os"
	"syscall"
	"unsafe"
)

// canvasSize returns the terminal columns, rows, and cursor aspect ratio
func canvasSize() (int, int, float64) {
	var size [4]uint16
	if _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(os.Stdout.Fd()), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&size)), 0, 0, 0); err != 0 {
		panic(err)
	}
	rows, cols, width, height := size[0], size[1], size[2], size[3]

	var whratio = defaultRatio
	if width > 0 && height > 0 {
		whratio = float64(height/rows) / float64(width/cols)
	}

	return int(cols), int(rows), whratio
}
