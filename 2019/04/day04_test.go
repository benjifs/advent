
package main

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	input := []int{111111, 223450, 123789}
	expected := []bool{true, false, false}
	
	for i, password := range input {
		if expected[i] != isValid(password) {
			t.Errorf("%d - expected %v", password, expected[i])
		}
	}
}

func TestHasDoubles(t *testing.T) {
	input := []int{112233, 123444, 111122}
	expected := []bool{true, false, true}

	for i, password := range input {
		if expected[i] != hasDoubles(password) {
			t.Errorf("%d - expected %v", password, expected[i])
		}
	}
}

