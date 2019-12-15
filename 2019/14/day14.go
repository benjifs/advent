package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"math"
)

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	reactions := parseInput(input)
	fmt.Println("pt1:", getMinimumOre(reactions))
}

func getMinimumOre(reactions map[string]Reaction) (int) {
	extra := map[string]int{}
	need := map[string]int{}
	need["FUEL"] = 1

	for len(need) > 1 || need["FUEL"] != 0 {
		for key, value := range need {
			if key == "ORE" {
				continue
			}
			reaction := reactions[key]
			n := int(math.Ceil(float64(value - extra[key]) / float64(reaction.result.quantity)))
			for _, value := range reaction.combinations {
				need[value.name] += n * value.quantity
			}
			extra[key] += n * reaction.result.quantity - value
			delete(need, key)
		}
	}

	return need["ORE"]
}

type Chemical struct {
	quantity int
	name string
}

type Reaction struct {
	result Chemical
	combinations []Chemical
}

func parseInput(input []string) (map[string]Reaction) {
	reactions := make(map[string]Reaction)

	for _, line := range input {
		reaction := parseReaction(line)
		reactions[reaction.result.name] = reaction
	}

	return reactions
}

func parseReaction(line string) (reaction Reaction) {
	split := strings.Split(line, " => ")
	combinations := strings.Split(split[0], ", ")

	chemicals := []Chemical{}
	for _, combination := range combinations {
		chemicals = append(chemicals, parseChemical(combination))
	}
	result := parseChemical(split[1])
	return Reaction{result, chemicals}
}

func parseChem(chem string) (string, int) {
	chemical := strings.Split(chem, " ")
	quantity, err := strconv.Atoi(chemical[0])
	if err != nil {
		panic(err)
	}
	return chemical[1], quantity
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

