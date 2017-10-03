package pxl

import "testing"

func TestMax(t *testing.T) {
	var want float64 = 5
	have := max(1, 2, 3, 4, 5)
	if want != have {
		t.Errorf("Max number mismatch. Want: %v, Have: %v\n", want, have)
	}
}

func TestScale(t *testing.T) {
	var whratio float64
	var want float64
	// 4x4 image fitting in a 2x2 terminal
	imgW, imgH, termW, termH := 4, 4, 2, 2
	whratio = 1

	want = 2
	have := scale(imgW, imgH, termW, termH, whratio)
	if want != have {
		t.Errorf("Image scale mismatch. Want: %v, Have: %v\n", want, have)
	}

	// 2x2 image fitting in a 4x4 terminal
	imgW, imgH, termW, termH = 2, 2, 4, 4
	whratio = 1

	want = 1
	have = scale(imgW, imgH, termW, termH, whratio)
	if want != have {
		t.Errorf("Image scale mismatch. Want: %v, Have: %v\n", want, have)
	}

	// 4x4 image fitting in a 2x1 terminal, with whratio = 2
	imgW, imgH, termW, termH = 4, 4, 2, 1
	whratio = 2

	want = 2
	have = scale(imgW, imgH, termW, termH, whratio)
	if want != have {
		t.Errorf("Image scale mismatch. Want: %v, Have: %v\n", want, have)
	}

}
