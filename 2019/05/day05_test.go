package main

import (
	"testing"
)

func TestSliceInput(t *testing.T) {
	input := "1,0,0,0,99"
	output := []int{1, 0, 0, 0, 99}
	actual, _ := inputToInt(input)
	if len(output) != len(actual) {
		t.Error("parsed input length does not match")
	}
	for i := 0; i < len(output); i++ {
		if output[i] != actual[i] {
			t.Error("parsed input does not match")
		}
	}
}

func TestGetOP(t *testing.T) {
	input := []int{1102, 99, 2, 0001}
	output := []int{2, 99, 2, 1}

	for i, in := range input {
		op := getOP(in)
		if op[0] != output[i] {
			t.Errorf("Expected: %d - Actual: %d\n", output[i], op[0])
		}
	}
}
