package main

import (
	"sync"
	"syscall"
	"unsafe"

	"github.com/nsf/termbox-go"
)

// canvasSize returns the terminal columns, rows, and cursor aspect ratio
func canvasSize() (int, int, float64) {
	cols, rows := termbox.Size()
	var whratio = defaultRatio
	if width, height, err := primarySize(); err == nil && width > 0 && height > 0 {
		whratio = float64(height/rows) / float64(width/cols)
	}
	return cols, rows, whratio
}

// Extracted (and trimmed down) from
// https://gist.github.com/dchapes/b90e4e0bde7d2672647b0d479ed42dbe
// (That's also why there are several functions and a bit more
//  boiler plate that required for this one user32.dll call).

const spiGetWorkArea = 0x0030

var winapi struct {
	once                 sync.Once
	systemParametersInfo uintptr
	err                  error
}

func loadWinAPI() {
	winapi.once.Do(func() {
		libuser32, err := syscall.LoadLibrary("user32.dll")
		if err != nil {
			winapi.err = err
			return
		}
		winapi.systemParametersInfo, winapi.err = syscall.GetProcAddress(libuser32, "SystemParametersInfoW")
	})
}

func systemParametersInfo(uiAction, uiParam uint32, pvParam unsafe.Pointer, fWinIni uint32) error {
	loadWinAPI()
	if winapi.err != nil {
		return winapi.err
	}

	r1, _, err := syscall.Syscall6(winapi.systemParametersInfo, 4,
		uintptr(uiAction), uintptr(uiParam),
		uintptr(pvParam), uintptr(fWinIni), 0, 0)
	if r1 == 0 {
		return err
	}
	return nil
}

// primarySize returns the size of the primary display monitor.
func primarySize() (width, height int, err error) {
	var rect struct {
		left   uint32
		top    uint32
		right  uint32
		bottom uint32
		// extra size as a safety margin
		_ uint32
		_ uint32
		_ uint32
		_ uint32
	}
	err = systemParametersInfo(spiGetWorkArea, 0, unsafe.Pointer(&rect.left), 0)
	if err != nil {
		return 0, 0, err
	}
	return int(rect.right) - int(rect.left),
		int(rect.bottom) - int(rect.top), nil
}
