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

func TestDay02_sample1(t *testing.T) {
	input := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	output := 3500
	actual, err := intCode(input)
	if err != nil {
		panic(err)
	}
	if output != actual[0] {
		t.Errorf("Expected: %d - Actual: %d", output, actual)
	}
}

func TestDay02_sample2(t *testing.T) {
	input := []int{1, 0, 0, 0, 99}
	output := 2
	actual, err := intCode(input)
	if err != nil {
		panic(err)
	}
	if output != actual[0] {
		t.Errorf("Expected: %d - Actual: %d", output, actual)
	}
}

func TestDay02_sample3(t *testing.T) {
	input := []int{2, 3, 0, 3, 99}
	output := 2
	actual, err := intCode(input)
	if err != nil {
		panic(err)
	}
	if output != actual[0] {
		t.Errorf("Expected: %d - Actual: %d", output, actual)
	}
}

func TestDay02_sample4(t *testing.T) {
	input := []int{2, 4, 4, 5, 99, 0}
	output := 2
	actual, err := intCode(input)
	if err != nil {
		panic(err)
	}
	if output != actual[0] {
		t.Errorf("Expected: %d - Actual: %d", output, actual)
	}
}

func TestDay02_sample5(t *testing.T) {
	input := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	output := 30
	actual, err := intCode(input)
	if err != nil {
		panic(err)
	}
	if output != actual[0] {
		t.Errorf("Expected: %d - Actual: %d", output, actual)
	}
}

