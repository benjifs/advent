package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
	"sort"
)

type asteroid struct {
	x, y int
	d int
	targets map[float64][]asteroid
}

func main() {
	lines, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	asteroids := initAsteroids(lines)
	best := calculateHits(asteroids)
	fmt.Printf("pt1: (%d, %d) -> %d\n", best.x, best.y, len(best.targets))

	destroyed := activateLaser(best)
	fmt.Println("pt2:", destroyed[199].x * 100 + destroyed[199].y)
}

func (a asteroid) distance(b asteroid) (int) {
	return int(math.Abs(float64(a.x - b.x)) + math.Abs(float64(a.y - b.y)))
}

func (a asteroid) angle(b asteroid) (float64) {
	angle := math.Atan2(float64(a.x - b.x), float64(b.y - a.y))
	angle = angle + math.Pi
	if angle == math.Pi * 2 {
		angle = 0
	}
	return angle
}

func activateLaser(a asteroid) (destroyed []asteroid) {
	var keys []float64
	for k := range a.targets {
		keys = append(keys, k)

		sort.Slice(a.targets[k], func(i, j int) (bool) {
			return a.targets[k][i].d < a.targets[k][j].d
		})
	}
	sort.Float64s(keys)

	for len(destroyed) != len(a.targets) {
		for _, k := range keys {
			if len(a.targets[k]) == 0 {
				delete(a.targets, k)
				continue
			}
			destroyed = append(destroyed, a.targets[k][0])
			a.targets[k] = a.targets[k][1:]
		}
	}
	return destroyed
}

func calculateHits(asteroids []asteroid) (best asteroid) {
	for i, a := range asteroids {
		targets := make(map[float64][]asteroid)
		for j, b := range asteroids {
			if i == j {
				continue
			}
			angle := a.angle(b)
			b.d = a.distance(b)
			targets[angle] = append(targets[angle], b)
		}
		a.targets = targets
		if len(best.targets) < len(a.targets) {
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

