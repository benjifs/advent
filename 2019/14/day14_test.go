package main

import (
	"testing"
)

func TestParseChemical(t *testing.T) {
	line := "3 ABC"
	chemical := parseChemical(line)

	if chemical.quantity != 3 && chemical.name != "ABC" {
		t.Errorf("Expected 3 ABC\n")
	}
}

func TestParseReaction(t *testing.T) {
	line := "1 A => 5 B"
	reaction := parseReaction(line)
	
	if len(reaction.combinations) != 1 {
		t.Errorf("Incorrect number of combinations\n")
	}

	line = "1 A, 2 C => 5 B"
	reaction = parseReaction(line)
	
	if len(reaction.combinations) != 2 {
		t.Errorf("Incorrect number of combinations\n")
	}
}

func TestParseInput(t *testing.T) {
	lines := []string{"10 ORE => 10 A", "1 ORE => 1 B", "7 A, 1 B => 1 C", "7 A, 1 C => 1 D", "7 A, 1 D => 1 E", "7 A, 1 E => 1 FUEL"}
	parsed := parseInput(lines)

	if len(parsed) != 6 {
		t.Errorf("Incorrect number of reactions\n")
	}
}

func TestPart1Sample1(t *testing.T) {
	lines := []string{"10 ORE => 10 A", "1 ORE => 1 B", "7 A, 1 B => 1 C", "7 A, 1 C => 1 D", "7 A, 1 D => 1 E", "7 A, 1 E => 1 FUEL"}
	expected := 31

	parsed := parseInput(lines)
	actual := getMinimumOre(parsed, 1)
	if expected != actual {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}

func TestPart2Input1(t *testing.T) {
	input, _ := readInput("test1.txt")
	reactions := parseInput(input)
	expected := 82892753

	actual := oreToFuel(reactions, 0, 100000000, 1000000000000)
	if expected != actual {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}

func TestPart2Input2(t *testing.T) {
	input, _ := readInput("test2.txt")
	reactions := parseInput(input)
	expected := 5586022

	actual := oreToFuel(reactions, 0, 100000000, 1000000000000)
	if expected != actual {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}

func TestPart2Input3(t *testing.T) {
	input, _ := readInput("test3.txt")
	reactions := parseInput(input)
	expected := 460664

	actual := oreToFuel(reactions, 0, 100000000, 1000000000000)
	if expected != actual {
		t.Errorf("Expected: %d - Actual: %d\n", expected, actual)
	}
}

