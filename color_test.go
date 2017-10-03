package pxl

import (
	"image/color"
	"testing"
)

func TestRGB(t *testing.T) {
	r, g, b := rgb(color.Black)
	var want = [3]uint16{0x00, 0x00, 0x00}
	var have = [3]uint16{r, g, b}
	if !eq(want, have) {
		t.Errorf("RGB data mismatch. Want: %v, Have: %v\n", want, have)
	}

	r, g, b = rgb(color.White)
	have = [3]uint16{r, g, b}
	want[0], want[1], want[2] = 0xff, 0xff, 0xff
	if !eq(want, have) {
		t.Errorf("RGB data mismatch. Want: %v, Have: %v\n", want, have)
	}

	r, g, b = rgb(color.RGBA{0xff, 0x00, 0x00, 0x00})
	have = [3]uint16{r, g, b}
	want[0], want[1], want[2] = 0xff, 0x00, 0x00
	if !eq(want, have) {
		t.Errorf("RGB data mismatch. Want: %v, Have: %v\n", want, have)
	}
}

func TestTermColor(t *testing.T) {
	want := uint16(16 + 1) // Black + default offset
	have := termColor(0, 0, 0)

	if want != have {
		t.Errorf("term color mismatch. Want: %v, Have: %v\n", want, have)
	}

	want = uint16(231 + 1) // White + default offset
	have = termColor(255, 255, 255)

	if want != have {
		t.Errorf("term color mismatch. Want: %v, Have: %v\n", want, have)
	}

	want = uint16(196 + 1) // Red + default offset
	have = termColor(255, 0, 0)

	if want != have {
		t.Errorf("term color mismatch. Want: %v, Have: %v\n", want, have)
	}
}

func eq(a, b [3]uint16) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}
