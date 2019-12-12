package main

import (
	"testing"
)

func TestPart1Sample1(t *testing.T) {
	moons := []Moon{}
	moons = append(moons, Moon{pos: Point{-1, 0, 2}})
	moons = append(moons, Moon{pos: Point{2, -10, -7}})
	moons = append(moons, Moon{pos: Point{4, -8, 8}})
	moons = append(moons, Moon{pos: Point{3, 5, -1}})

	expected := 179
	actual := runNSimulations(moons, 10)
	if expected != actual {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}

func TestPart1Sample2(t *testing.T) {
	moons := []Moon{}
	moons = append(moons, Moon{pos: Point{-8, -10, 0}})
	moons = append(moons, Moon{pos: Point{5, 5, 10}})
	moons = append(moons, Moon{pos: Point{2, -7, 3}})
	moons = append(moons, Moon{pos: Point{9, -8, -3}})

	expected := 1940
	actual := runNSimulations(moons, 100)
	if expected != actual {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}

func TestPart2Sample1(t *testing.T) {
	moons := []Moon{}
	moons = append(moons, Moon{pos: Point{-1, 0, 2}})
	moons = append(moons, Moon{pos: Point{2, -10, -7}})
	moons = append(moons, Moon{pos: Point{4, -8, 8}})
	moons = append(moons, Moon{pos: Point{3, 5, -1}})

	expected := 2772
	actual := runSimulations(moons)
	if expected != actual {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}

func TestPart2Sample2(t *testing.T) {
	moons := []Moon{}
	moons = append(moons, Moon{pos: Point{-8, -10, 0}})
	moons = append(moons, Moon{pos: Point{5, 5, 10}})
	moons = append(moons, Moon{pos: Point{2, -7, 3}})
	moons = append(moons, Moon{pos: Point{9, -8, -3}})

	expected := 4686774924
	actual := runSimulations(moons)
	if expected != actual {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}

