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

	fmt.Println("pt1", runNSimulations(moons, 1000))
	fmt.Println("pt2", runSimulations(moons))
}

func runSimulations(moons []Moon) (int) {
	xs := make(map[string]bool)
	ys := make(map[string]bool)
	zs := make(map[string]bool)
	matched := [3]int{0, 0, 0}

	for i := 0; matched[0] == 0 || matched[1] == 0 || matched[2] == 0; i++ {
		keys := getKeys(moons)

		if matched[0] == 0 && xs[keys[0]] {
			matched[0] = i
		}
		if matched[1] == 0 && ys[keys[1]] {
			matched[1] = i
		}
		if matched[2] == 0 && zs[keys[2]] {
			matched[2] = i
		}
		xs[keys[0]] = true
		ys[keys[1]] = true
		zs[keys[2]] = true

		moons = simulate(moons)
	}
	return lcm(lcm(matched[0], matched[1]), matched[2])
}

func getKeys(moons []Moon) ([]string) {
	x := fmt.Sprintf("%d:%d %d:%d %d:%d %d:%d", moons[0].pos.x, moons[0].vel.x, moons[1].pos.x, moons[1].vel.x, moons[2].pos.x, moons[2].vel.x, moons[3].pos.x, moons[3].vel.x)
	y := fmt.Sprintf("%d:%d %d:%d %d:%d %d:%d", moons[0].pos.y, moons[0].vel.y, moons[1].pos.y, moons[1].vel.y, moons[2].pos.y, moons[2].vel.y, moons[3].pos.y, moons[3].vel.y)
	z := fmt.Sprintf("%d:%d %d:%d %d:%d %d:%d", moons[0].pos.z, moons[0].vel.z, moons[1].pos.z, moons[1].vel.z, moons[2].pos.z, moons[2].vel.z, moons[3].pos.z, moons[3].vel.z)
	return []string{x, y, z}
}

func runNSimulations(moons []Moon, n int) (int) {
	for i := 1; i <= n; i++ {
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

func (m Moon) getTotalEnergy() (int) {
	return m.pos.getEnergy() * m.vel.getEnergy()
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

func gcd(a, b int) (int) {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}

func lcm(a, b int) (int) {
	return a / gcd(a, b) * b
}
