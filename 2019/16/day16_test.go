package main

import (
	"testing"
)

func TestGetNDigits(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expected := []string{"1234", "123", "123456"}

	for _, val := range expected {
		actual := getNDigits(input, len(val))
		if val != actual {
			t.Errorf("Expected: %v - Actual: %v\n", val, actual)
		}
	}
}

func TestPart1SinglePhase(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expected := []int{4, 8, 2, 2, 6, 1, 5, 8}

	actual := doPhase(input, []int{0, 1, 0, -1})
	for i := range actual {
		if expected[i] != actual[i] {
			t.Errorf("Digit %d does not match\n", i)
		}
	}
}

func TestPart1NPhases(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expected := "01029498"

	phases := doNPhases(input, []int{0, 1, 0, -1}, 4)
	actual := getNDigits(phases, 8)
	if expected != actual {
		t.Errorf("Expected: %v - Actual: %v\n", expected, actual)
	}
}

