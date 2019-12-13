package main

import (
	"testing"
)

func TestCountBlocks(t *testing.T) {
	grid := make(map[Point]int)
	grid[Point{1, 2}] = 3
	grid[Point{6, 5}] = 4
	grid[Point{0, 0}] = 1
	grid[Point{1, 0}] = 2
	grid[Point{2, 0}] = 3
	grid[Point{0, 2}] = 2

	game := Game{screen: grid}
	expected := 2

	actual := game.countBlocks()
	if expected != actual {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}

