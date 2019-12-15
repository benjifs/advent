package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {

	fmt.Println("pt1")
}

type Chemical struct {
	quantity int
	name string
}

type Reaction struct {
	result Chemical
	combination []Chemical
}

func parseInput(input []string) ([]Reaction) {
	reactions := []Reaction{}

	for _, line := range input {
		reactions = append(reactions, parseReaction(line))
	}

	return reactions
}

func parseReaction(line string) (Reaction) {
	split := strings.Split(line, " => ")
	combinations := strings.Split(split[0], ", ")

	chemicals := []Chemical{}
	for _, combination := range combinations {
		chemicals = append(chemicals, parseChemical(combination))
	}
	result := parseChemical(split[1])

	return Reaction{result, chemicals}
}

func parseChemical(chem string) (Chemical) {
	parsed := strings.Split(chem, " ")
	quantity, err := strconv.Atoi(parsed[0])
	if err != nil {
		panic(err)
	}

	return Chemical{quantity, parsed[1]}
}

func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input, nil
}

func inputToInt(input string) ([]int, error) {
	var out []int
	for _, val := range strings.Split(input, ",") {
		num, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		out = append(out, num)
	}
	return out, nil
}

