package main

import (
	"testing"
)

func TestPart1Sample1(t *testing.T) {
	line := "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"

	input, _ := inputToInt(line)
	actual := intCode(Amplifier{mem: input})

	for i, val := range actual.mem {
		if val != input[i] {
			t.Errorf("Program does not output itself")
		}
	}
}

func TestPart1Sample2(t *testing.T) {
	line := "1102,34915192,34915192,7,4,7,99,0"
	expected := 16

	input, _ := inputToInt(line)
	actual := intCode(Amplifier{mem: input})

	signal := actual.signal
	digits := 0
	for signal > 0 {
		signal = signal / 10
		digits++
	}
	if expected != digits {
		t.Errorf("Expected: %d - Actual: %d\n", expected, digits)
		t.Errorf("%v\n", actual)
	}
}

func TestPart1Sample3(t *testing.T) {
	line := "104,1125899906842624,99"
	expected := 1125899906842624

	input, _ := inputToInt(line)
	actual := intCode(Amplifier{mem: input})
	if expected != actual.signal {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual.signal)
	}
}

