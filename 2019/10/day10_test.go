package main

import (
	"testing"
)

func TestInitAsteroids(t *testing.T) {
	lines := []string{".#..#", ".....", "#####", "....#", "...##"}
	expected := 10

	asteroids := initAsteroids(lines)
	actual := len(asteroids)

	if expected != actual {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}

func TestPart1Sample1(t *testing.T) {
	lines := []string{".#..#", ".....", "#####", "....#", "...##"}
	expected := 8

	asteroids := initAsteroids(lines)
	actual := calculateHits(asteroids)

	if expected != actual.count {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}

func TestPart1Sample2(t *testing.T) {
	lines := []string{"......#.#.", "#..#.#....", "..#######.", ".#.#.###..", ".#..#.....", "..#....#.#", "#..#....#.", ".##.#..###", "##...#..#.", ".#....####"}
	expected := 33

	asteroids := initAsteroids(lines)
	actual := calculateHits(asteroids)

	if expected != actual.count {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}
