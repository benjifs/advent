package main

import (
	"testing"
)

func TestGetPhaseSettings(t *testing.T) {
	seq := []int{0, 1, 2, 3, 4}
	expected := 120
	
	actual := len(getPhaseSettings(seq))
	if expected != actual {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}

func TestPart1Sample1(t *testing.T) {
	line := "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"
	sequence := []int{4, 3, 2, 1, 0}
	expected := 43210

	input, _ := inputToInt(line)
	actual := runWithPhaseSetting(input, sequence)
	if expected != actual {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}

func TestPart1Sample2(t *testing.T) {
	line := "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0"
	sequence := []int{0, 1, 2, 3, 4}
	expected := 54321

	input, _ := inputToInt(line)
	actual := runWithPhaseSetting(input, sequence)
	if expected != actual {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}

func TestPart1Sample3(t *testing.T) {
	line := "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0"
	sequence := []int{1, 0, 4, 3, 2}
	expected := 65210

	input, _ := inputToInt(line)
	actual := runWithPhaseSetting(input, sequence)
	if expected != actual {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}

