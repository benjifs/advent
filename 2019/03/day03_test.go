package main

import (
	"testing"
)

func TestPart01_sample1(t *testing.T) {
	wire1 := []string{"R8", "U5", "L5", "D3"}
	wire2 := []string{"U7", "R6", "D4", "L4"}
	expected := 6

	grid1 := createMap(wire1)
	grid2 := createMap(wire2)

	actual := calculateDistance(grid1.points, grid2.points)
	if actual != expected {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}

func TestPart01_sample2(t *testing.T) {
	wire1 := []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
	wire2 := []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}
	expected := 159

	grid1 := createMap(wire1)
	grid2 := createMap(wire2)

	actual := calculateDistance(grid1.points, grid2.points)
	if actual != expected {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}

func TestPart01_sample3(t *testing.T) {
	wire1 := []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}
	wire2 := []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}
	expected := 135

	grid1 := createMap(wire1)
	grid2 := createMap(wire2)

	actual := calculateDistance(grid1.points, grid2.points)
	if actual != expected {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}

func TestPart02_sample1(t *testing.T) {
	wire1 := []string{"R8", "U5", "L5", "D3"}
	wire2 := []string{"U7", "R6", "D4", "L4"}
	expected := 30

	grid1 := createMap(wire1)
	grid2 := createMap(wire2)

	actual := calculateSteps(grid1.points, grid2.points)
	if actual != expected {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}

func TestPart02_sample2(t *testing.T) {
	wire1 := []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
	wire2 := []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}
	expected := 610

	grid1 := createMap(wire1)
	grid2 := createMap(wire2)

	actual := calculateSteps(grid1.points, grid2.points)
	if actual != expected {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}

func TestPart03_sample3(t *testing.T) {
	wire1 := []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}
	wire2 := []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}
	expected := 410

	grid1 := createMap(wire1)
	grid2 := createMap(wire2)

	actual := calculateSteps(grid1.points, grid2.points)
	if actual != expected {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}

