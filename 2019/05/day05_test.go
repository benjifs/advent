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

func TestPt2_input1(t *testing.T) {
	line := "3,9,8,9,10,9,4,9,99,-1,8"
	instructions, _ := inputToInt(line)
	inputs := []int{8, 5, 3}
	outputs := []int{1, 0, 0}

	for i, input := range inputs {
		_, actual := intCode(instructions, input)
		if actual != outputs[i] {
			t.Errorf("Expected: %d - Actual: %d\n", outputs[i], actual)
		}
	}
}

func TestPt2_input2(t *testing.T) {
	line := "3,9,7,9,10,9,4,9,99,-1,8"
	instructions, _ := inputToInt(line)
	inputs := []int{8, 5, 9}
	outputs := []int{0, 1, 0}

	for i, input := range inputs {
		_, actual := intCode(instructions, input)
		if actual != outputs[i] {
			t.Errorf("Expected: %d - Actual: %d\n", outputs[i], actual)
		}
	}
}

func TestPt2_input3(t *testing.T) {
	line := "3,3,1108,-1,8,3,4,3,99"
	instructions, _ := inputToInt(line)
	inputs := []int{8, 5, 3}
	outputs := []int{1, 0, 0}

	for i, input := range inputs {
		_, actual := intCode(instructions, input)
		if actual != outputs[i] {
			t.Errorf("Expected: %d - Actual: %d\n", outputs[i], actual)
		}
	}
}

func TestPt2_input4(t *testing.T) {
	line := "3,3,1107,-1,8,3,4,3,99"
	instructions, _ := inputToInt(line)
	inputs := []int{8, 5, 9}
	outputs := []int{0, 1, 0}

	for i, input := range inputs {
		_, actual := intCode(instructions, input)
		if actual != outputs[i] {
			t.Errorf("Expected: %d - Actual: %d\n", outputs[i], actual)
		}
	}
}

func TestPt2_input5(t *testing.T) {
	line := "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9"
	instructions, _ := inputToInt(line)
	inputs := []int{0, 5, 9}
	outputs := []int{0, 1, 1}

	for i, input := range inputs {
		_, actual := intCode(instructions, input)
		if actual != outputs[i] {
			t.Errorf("Expected: %d - Actual: %d\n", outputs[i], actual)
		}
	}
}

func TestPt2_input6(t *testing.T) {
	line := "3,3,1105,-1,9,1101,0,0,12,4,12,99,1"
	instructions, _ := inputToInt(line)
	inputs := []int{0, 5, 9}
	outputs := []int{0, 1, 1}

	for i, input := range inputs {
		_, actual := intCode(instructions, input)
		if actual != outputs[i] {
			t.Errorf("Expected: %d - Actual: %d\n", outputs[i], actual)
		}
	}
}

func TestPt2_input7(t *testing.T) {
	line := "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"
	instructions, _ := inputToInt(line)
	inputs := []int{2, 8, 10}
	outputs := []int{999, 1000, 1001}

	for i, input := range inputs {
		_, actual := intCode(instructions, input)
		if actual != outputs[i] {
			t.Errorf("Expected: %d - Actual: %d\n", outputs[i], actual)
		}
	}
}
