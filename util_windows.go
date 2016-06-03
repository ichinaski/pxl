// +build windows

package main

import (
	"os"
	"syscall"
	"unsafe"
)

type short int16
type word uint16

type coord struct {
	x short
	y short
}

type smallRect struct {
	left   short
	top    short
	right  short
	bottom short
}

type consoleScreenBufferInfo struct {
	size              coord
	cursorPosition    coord
	attributes        word
	window            smallRect
	maximumWindowSize coord
}

type consoleFontInfo struct {
	nfont    uint32
	fontSize coord
}

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	procGetConsoleScreenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")
	procGetCurrentConsoleFont      = kernel32.NewProc("GetCurrentConsoleFont")
)

// canvasSize returns the terminal columns, rows, and cursor aspect ratio
func canvasSize() (int, int, float64) {
	var csbi consoleScreenBufferInfo
	r1, _, err := procGetConsoleScreenBufferInfo.Call(os.Stdout.Fd(), uintptr(unsafe.Pointer(&csbi)))
	if r1 == 0 {
		panic(err)
	}
	var cfi consoleFontInfo
	r1, _, err = procGetCurrentConsoleFont.Call(os.Stdout.Fd(), uintptr(0), uintptr(unsafe.Pointer(&cfi)))
	if r1 == 0 {
		panic(err)
	}
	cols, rows := csbi.window.right-csbi.window.left, csbi.window.bottom-csbi.window.top
	width, height := cfi.fontSize.x*cols, cfi.fontSize.y*rows
	return int(cols), int(rows), float64(height/rows) / float64(width/cols)
}
