package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lines, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	orbitMap := makeOrbitMap(lines)

	fmt.Println("pt1:", numOrbits(orbitMap))
	fmt.Println("pt2:", getClosestPath(getRoute("YOU", orbitMap), getRoute("SAN", orbitMap)))
}

func makeOrbitMap(lines []string) (map[string]string) {
	orbitMap := make(map[string]string)
	for _, line := range lines {
		vals := strings.Split(line, ")")
		orbitMap[vals[1]] = vals[0]
	}
	return orbitMap
}

func getRoute(start string, orbitMap map[string]string) ([]string) {
	var route []string

	loc := orbitMap[start]
	for {
		if loc == "" || loc == "COM" {
			break
		}
		route = append(route, loc)
		loc = orbitMap[loc]
	}

	return route
}

func getClosestPath(path1, path2 []string) (int) {
	for i, val1 := range path1 {
		for j, val2 := range path2 {
			if val1 == val2 {
				return i + j
			}
		}
	}
	return 0
}

func numOrbits(orbitMap map[string]string) (int) {
	distance := 0
	for k, _ := range orbitMap {
		distance += getOrbits(k, orbitMap)
	}
	return distance
}

func getOrbits(loc string, orbitMap map[string]string) (int) {
	if orbitMap[loc] == "" {
		return 0
	}
	return getOrbits(orbitMap[loc], orbitMap) + 1
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

