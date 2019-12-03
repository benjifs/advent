package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
const MaxInt = int(^uint(0) >> 1)

func main() {
	lines, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	path1 := createCoords(strings.Split(lines[0], ","))
	path2 := createCoords(strings.Split(lines[1], ","))

	intersects := calculateIntersects(path1, path2)

	fmt.Printf("distance: %d\n", getShortestDistance(intersects))
	fmt.Printf("steps: %d\n", getBestSteps(intersects))
}

func createCoords(input []string) ([]Point) {
	origin := Point{x: 0, y: 0, steps: 0}
	coords := []Point{origin}

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
		origin.x += offset.x * distance
		origin.y += offset.y * distance
		origin.steps += distance

		coords = append(coords, origin)
	}
	return coords
}

func getShortestDistance(points []Point) (int) {
	origin := Point{x: 0, y: 0}
	manhattan := MaxInt
	for _, point := range points {
		distance := abs(origin.x - point.x) + abs(origin.y - point.y)
		if distance != 0 && distance < manhattan {
			manhattan = distance
		}
	}
	return manhattan
}

func getBestSteps(points []Point) (int) {
	steps := MaxInt
	for _, point := range points {
		if point.steps != 0 && point.steps < steps {
			steps = point.steps
		}
	}
	return steps
}

func makeLine(p1, p2 Point) (Line) {
	line := Line{start: p1, end: p2}

	line.a = p1.y - p2.y
	line.b = p2.x - p1.x
	line.c = -(p1.x * p2.y - p2.x * p1.y)

	return line
}

func calculateIntersects(path1 []Point, path2 []Point) ([]Point) {
	var intersects []Point

	for i := 0; i < len(path1) - 1; i++ {
		for j := 0; j < len(path2) - 1; j++ {
			line1 := makeLine(path1[i], path1[i + 1])
			line2 := makeLine(path2[j], path2[j + 1])

			intersect, found := findIntersect(line1, line2)
			if found {
				intersects = append(intersects, intersect)
			}
		}
	}
	return intersects
}

func inLine(point Point, line Line) (bool) {
	var minX, maxX, minY, maxY int

	if line.start.x < line.end.x {
		minX = line.start.x
		maxX = line.end.x
	} else {
		maxX = line.start.x
		minX = line.end.x
	}
	if line.start.y < line.end.y {
		minY = line.start.y
		maxY = line.end.y
	} else {
		maxY = line.start.y
		minY = line.end.y
	}

	return minX <= point.x && point.x <= maxX && minY <= point.y && point.y <= maxY
}

func calculateSteps(intersect, p1, p2 Point) (int) {
	steps := p1.steps + p2.steps

	steps -= abs(p1.x - intersect.x)
	steps -= abs(p2.x - intersect.x)
	steps -= abs(p1.y - intersect.y)
	steps -= abs(p2.y - intersect.y)

	return steps
}

func findIntersect(l1, l2 Line) (Point, bool) {
	d := l1.a * l2.b - l1.b * l2.a
	dx := l1.c * l2.b - l1.b * l2.c
	dy := l1.a * l2.c - l1.c * l2.a

	if d != 0 {
		x := dx / d
		y := dy / d
		p := Point{x: x, y: y}
		if inLine(p, l1) && inLine(p, l2) {
			p.steps = calculateSteps(p, l1.end, l2.end)
			return p, true
		}
	}
	return Point{}, false
}


type Point struct {
	x, y, steps int
}

type Line struct {
	a, b, c int
	start, end Point
}

func abs(x int) (int) {
	if x < 0 {
		return x * -1
	}
	return x
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
