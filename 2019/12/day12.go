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

	moons := []Moon{}
	for _, line := range lines {
		moons = append(moons, parseMoon(line))
	}

	fmt.Println("pt1", runSimulation(moons, 1000))
}

func runSimulation(moons []Moon, steps int) (int) {
	for i := 1; i <= steps; i++ {
		moons = simulate(moons)
	}
	total := 0
	for _, moon := range moons {
		total += moon.getTotalEnergy()
	}
	return total
}

func simulate(moons []Moon) ([]Moon) {
	for i, moon := range moons {
		for j, m := range moons {
			if i == j {
				continue
			}
			if moon.pos.x < m.pos.x {
				moon.vel.x++
			} else if moon.pos.x > m.pos.x {
				moon.vel.x--
			}
			if moon.pos.y < m.pos.y {
				moon.vel.y++
			} else if moon.pos.y > m.pos.y {
				moon.vel.y--
			}
			if moon.pos.z < m.pos.z {
				moon.vel.z++
			} else if moon.pos.z > m.pos.z {
				moon.vel.z--
			}
			moons[i] = moon
		}
	}
	for i := range moons {
		moons[i].addVelocity()
	}
	return moons
}

type Point struct {
	x, y, z int
}

func (p Point) getEnergy() (int) {
	return abs(p.x) + abs(p.y) + abs(p.z)
}

type Moon struct {
	pos Point
	vel Point
}

func (m Moon) getPotentialEnergy() (int) {
	return m.pos.getEnergy()
}

func (m Moon) getKineticEnergy() (int) {
	return m.vel.getEnergy()
}

func (m Moon) getTotalEnergy() (int) {
	return m.getPotentialEnergy() * m.getKineticEnergy()
}

func (m *Moon) addVelocity() {
	m.pos.x += m.vel.x
	m.pos.y += m.vel.y
	m.pos.z += m.vel.z
}

func parseMoon(line string) (Moon) {
	// remove <>
	line = line[1:len(line) - 1]
	// separate variables
	vars := strings.Split(line, ", ")

	for i := range vars {
		tmp := strings.Split(vars[i], "=")
		vars[i] = tmp[1]
	}

	x, _ := strconv.Atoi(vars[0])
	y, _ := strconv.Atoi(vars[1])
	z, _ := strconv.Atoi(vars[2])

	return Moon{pos: Point{x, y, z}}
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

func abs(a int) (int) {
	if a < 0 {
		return a * -1
	}
	return a
}

