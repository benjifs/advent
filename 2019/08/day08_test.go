package main

import (
	"testing"
)

func TestCheckLayer(t *testing.T) {
	line := "123456789012"
	width, height := 3, 2
	expected := 1

	input, _ := inputToInt(line)
	img := generateImage(input, width, height)

	actual := checkLayer(img)
	if expected != actual {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}
