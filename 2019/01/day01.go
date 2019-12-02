package main

import (  
	"bufio"
	"fmt"
	"os"
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

func part1(lines []int) (int) {
	var total = 0
	for _, mass := range lines {
		total += (mass / 3) - 2
	}
	return total
}

func part2(lines []int) (int) {
	var total = 0
	for _, mass := range lines {
		for mass > 0 {
			mass = (mass / 3) - 2
			if mass > 0 {
				total += mass
			}
		}
	}
	return total
}

func readLines(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		lines = append(lines, i)
	}
	return lines, nil
}

