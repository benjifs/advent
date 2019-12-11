package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
)

type asteroid struct {
	x, y int
	count int
}

func main() {
	lines, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	asteroids := initAsteroids(lines)
	best := calculateHits(asteroids)
	fmt.Printf("pt1: (%d, %d) -> %d\n", best.x, best.y, best.count)
}

func (a asteroid) distance(b asteroid) (int) {
	return int(math.Abs(float64(a.x - b.x)) + math.Abs(float64(a.y - b.y)))
}

func (a asteroid) angle(b asteroid) float64 {
	return math.Atan2(float64(a.x - b.x), float64(a.y - b.y))
}

func calculateHits(asteroids []asteroid) (best asteroid) {
	for i, a := range asteroids {
		sight := make(map[float64]int)
		for j, b := range asteroids {
			if i == j {
				continue
			}
			angle := a.angle(b)
			distance := a.distance(b)

			if val, ok := sight[angle]; ok {
				if val > distance {
					sight[angle] = distance
				}
			} else {
				sight[angle] = distance
			}
		}
		a.count = len(sight)
		if best.count < a.count {
			best = a
		}
	}
	return best
}

func initAsteroids(lines []string) (asteroids []asteroid) {
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				asteroids = append(asteroids, asteroid{x: x, y: y})
			}
		}
	}
	return asteroids
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

