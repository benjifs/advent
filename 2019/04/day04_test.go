
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

