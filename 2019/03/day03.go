package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	lines, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	var grids []Grid
	for _, line := range lines {
		input := strings.Split(line, ",")
		grid := createMap(input)
		grids = append(grids, grid)
	}

	// fmt.Printf("distance: %d\n", calculateDistance(grids[0].points, grids[1].points))
	fmt.Printf("distance: %d\n", calculateSteps(grids[0].points, grids[1].points))
}

func calculateDistance(path1 []Point, path2 []Point) (int) {
	manhattan := 999
	origin := Point{x: 0, y: 0}
	for _, point1 := range path1 {
		for _, point2 := range path2 {
			if point1.x == point2.x && point1.y == point2.y {
				tmp := abs(origin.x - point1.x) + abs(origin.y - point1.y)
				if tmp != 0 && tmp < manhattan {
					manhattan = tmp
				}
			}
		}
	}
	return manhattan
}

func calculateSteps(path1 []Point, path2 []Point) (int) {
	steps := 99999
	for i, point1 := range path1 {
		for j, point2 := range path2 {
			if point1.x == point2.x && point1.y == point2.y {
				tmp := i + j
				if tmp != 0 && tmp < steps {
					steps = tmp
				}
			}
		}
	}
	return steps
}

func abs(x int) (int) {
	if x < 0 {
		return x * -1
	}
	return x
}

type Point struct {
	x int
	y int
}

type Grid struct {
	points []Point
}

func createMap(input []string) (Grid) {
	origin := Point{x: 0, y: 0}
	var grid Grid

	grid.points = append(grid.points, origin)

	for _, val := range input {
		offset := Point{x: 0, y: 0}
		switch val[0:1] {
			case "U":
				offset.y = 1
			case "D":
				offset.y = -1
			case "R":
				offset.x = 1
			case "L":
				offset.x = -1
		}
		distance, err := strconv.Atoi(val[1:])
		if err != nil {
			panic(err)
		}
		for i := 0; i < distance; i++ {
			origin.x += offset.x
			origin.y += offset.y

			grid.points = append(grid.points, origin)
		}
	}
	return grid
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

