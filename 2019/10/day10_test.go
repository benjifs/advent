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

	if expected != len(actual.targets) {
		t.Errorf("Expected: %d - Actual: %d\n", expected, len(actual.targets))
	}
}

func TestPart1Sample2(t *testing.T) {
	lines := []string{"......#.#.", "#..#.#....", "..#######.", ".#.#.###..", ".#..#.....", "..#....#.#", "#..#....#.", ".##.#..###", "##...#..#.", ".#....####"}
	expected := 33

	asteroids := initAsteroids(lines)
	actual := calculateHits(asteroids)

	if expected != len(actual.targets) {
		t.Errorf("Expected: %d - Actual: %d\n", expected, len(actual.targets))
	}
}

func TestSampleFilePt1(t *testing.T) {
	lines, err := readInput("test.txt")
	if err != nil {
		panic(err)
	}
	expected := 210

	asteroids := initAsteroids(lines)
	best := calculateHits(asteroids)

	if expected != len(best.targets) {
		t.Errorf("Expected: %d - Actual: %d\n", expected, len(best.targets))
	}
}

func TestSampleFilePt2(t *testing.T) {
	lines, err := readInput("test.txt")
	if err != nil {
		panic(err)
	}
	asteroids := initAsteroids(lines)
	best := calculateHits(asteroids)
	destroyed := activateLaser(best)

	expected := 802
	actual := destroyed[199].x * 100 + destroyed[199].y

	if expected != actual {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}

