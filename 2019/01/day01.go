package main

import (  
	"bufio"
	"fmt"
	"os"
	"math"
	"strconv"
)

func main() {
	var lines, err = readLines("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("part 1:", part1(lines))
	fmt.Println("part 2:", part2(lines))
}

func part1(lines []string) (int) {
	var total = 0
	for _, line := range lines {
		var mass, _ = strconv.ParseFloat(line, 64)

		total = total + int(math.Floor(mass / 3)) - 2
	}
	return total
}

func part2(lines []string) (int) {
	var total = 0
	for _, line := range lines {
		var mass, _ = strconv.ParseFloat(line, 64)

		for mass > 0 {
			mass = math.Floor(mass / 3) - 2
			if mass > 0 {
				total = total + int(mass)
			}
		}
	}
	return total
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

