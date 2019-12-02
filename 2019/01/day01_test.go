package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	inputs := []int{12, 14, 1969, 100756}
	outputs := []int{2, 2, 654, 33583}

	for i := 0; i < len(inputs); i++ {
		result := part1(append([]int{}, inputs[i]))
		if result != outputs[i] {
			t.Errorf("Input: %d -- Expected: %d - Actual: %d", inputs[i], outputs[i], result)
		}
	}
}

func TestPartTwo(t *testing.T) {
	inputs := []int{12, 14, 1969, 100756}
	outputs := []int{2, 2, 966, 50346}

	for i := 0; i < len(inputs); i++ {
		result := part2(append([]int{}, inputs[i]))
		if result != outputs[i] {
			t.Errorf("Input: %d -- Expected: %d - Actual: %d", inputs[i], outputs[i], result)
		}
	}
}

