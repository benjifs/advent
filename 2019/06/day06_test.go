package main

import (
	"testing"
)

func TestMakeOrbitMap(t *testing.T) {
	input := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}

	orbitMap := makeOrbitMap(input)

	if orbitMap["B"] != "COM" {
		t.Error("There was an error making the Orbit Map")
	}
	if orbitMap["C"] != "B" {
		t.Error("There was an error making the Orbit Map")
	}
	if orbitMap["COM"] != "" {
		t.Error("There was an error making the Orbit Map")
	}
}

func TestPt1(t *testing.T) {
	input := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}
	expected := 42

	orbitMap := makeOrbitMap(input)
	actual := numOrbits(orbitMap)
	if expected != actual {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}

func TestPt2(t *testing.T) {
	input := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"}
	expected := 4

	orbitMap := makeOrbitMap(input)
	actual := getClosestPath(getRoute("YOU", orbitMap), getRoute("SAN", orbitMap))
	if expected != actual {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}

func TestGetRoute(t *testing.T) {
	input := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"}
	expected := 6

	orbitMap := makeOrbitMap(input)
	actual := getRoute("YOU", orbitMap)
	if expected != len(actual) {
		t.Errorf("Expected: %d - Actual: %d\n", expected, len(actual))
	}
}

