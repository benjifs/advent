package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

type Point struct {
	x, y int
}

func main() {
	lines, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	input, _ := inputToInt(lines[0])
	for i := 0; i < 1000; i++ {
		input = append(input, 0)
	}

	fmt.Println("pt1:", len(paint(input, 0)))
	drawGrid(paint(input, 1))
}

func drawGrid(grid map[Point]int) {
	minX, maxX := MaxInt, MinInt
	minY, maxY := MaxInt, MinInt

	for point, _ := range grid {
		if point.x < minX {
			minX = point.x
		}
		if point.x > maxX {
			maxX = point.x
		}
		if point.y < minY {
			minY = point.y
		}
		if point.y > maxY {
			maxY = point.y
		}
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if grid[Point{x, y}] == 1 {
				fmt.Printf("█")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func paint(input []int, startColor int) (map[Point]int) {
	robot := 0
	point := Point{x: 0, y: 0}
	amp := Amplifier{mem: input}

	grid := make(map[Point]int)
	grid[point] = startColor
	for {
		if amp.halt {
			break
		}

		signal := 0
		if c, ok := grid[point]; ok {
			signal = c
		}

		amp.input = append(amp.input, signal)
		amp = intCode(amp)
		color := amp.signal
		amp = intCode(amp)
		direction := amp.signal

		grid[point] = color

		if direction == 0 {
			robot--
			if robot < 0 {
				robot = 3
			}
		} else if direction == 1 {
			robot++
			if robot > 3 {
				robot = 0
			}
		}
		if robot == 0 {
			point.y--
		} else if robot == 1 {
			point.x++
		} else if robot == 2 {
			point.y++
		} else if robot == 3 {
			point.x--
		}
	}
	return grid
}

func getOP(code int) ([]int) {
	op := code % 100
	param1 := (code / 100) % 10
	param2 := (code / 1000) % 10
	param3 := (code / 10000) % 10

	return []int{op, param1, param2, param3}
}

func getParam(mem []int, index int, mode int, relativeBase int) (int) {
	if mode == 0 {
		return mem[index]
	}
	if mode == 2 {
		return mem[index] + relativeBase
	}
	return index
}

type Amplifier struct {
	mem []int
	position int
	signal int
	input []int
	halt bool
	relativeBase int
}

func intCode(amp Amplifier) (Amplifier) {
	i := amp.position
	mem := amp.mem
	input := amp.input
	signal := amp.signal
	relativeBase := amp.relativeBase

	for i < len(mem) {
		ops := getOP(mem[i])

		op := ops[0]
		switch op {
			case 99:
				return Amplifier{mem: mem, signal: signal, position: i, input: input, halt: true, relativeBase: relativeBase}
			case 1:
				param1 := getParam(mem, i + 1, ops[1], relativeBase)
				param2 := getParam(mem, i + 2, ops[2], relativeBase)
				out := getParam(mem, i + 3, ops[3], relativeBase)
				mem[out] = mem[param1] + mem[param2]
				i += 4
			case 2:
				param1 := getParam(mem, i + 1, ops[1], relativeBase)
				param2 := getParam(mem, i + 2, ops[2], relativeBase)
				out := getParam(mem, i + 3, ops[3], relativeBase)
				mem[out] = mem[param1] * mem[param2]
				i += 4
			case 3:
				out := getParam(mem, i + 1, ops[1], relativeBase)
				mem[out], input = input[0], input[1:]

				i += 2
			case 4:
				out := getParam(mem, i + 1, ops[1], relativeBase)
				signal = mem[out]

				return Amplifier{mem: mem, signal: signal, position: i + 2, input: input, relativeBase: relativeBase}
			case 5:
				param1 := getParam(mem, i + 1, ops[1], relativeBase)
				param2 := getParam(mem, i + 2, ops[2], relativeBase)

				if mem[param1] != 0 {
					i = mem[param2]
				} else {
					i += 3
				}
			case 6:
				param1 := getParam(mem, i + 1, ops[1], relativeBase)
				param2 := getParam(mem, i + 2, ops[2], relativeBase)

				if mem[param1] == 0 {
					i = mem[param2]
				} else {
					i += 3
				}
			case 7:
				param1 := getParam(mem, i + 1, ops[1], relativeBase)
				param2 := getParam(mem, i + 2, ops[2], relativeBase)
				out := getParam(mem, i + 3, ops[3], relativeBase)

				if mem[param1] < mem[param2] {
					mem[out] = 1
				} else {
					mem[out] = 0
				}
				i += 4
			case 8:
				param1 := getParam(mem, i + 1, ops[1], relativeBase)
				param2 := getParam(mem, i + 2, ops[2], relativeBase)
				out := getParam(mem, i + 3, ops[3], relativeBase)

				if mem[param1] == mem[param2] {
					mem[out] = 1
				} else {
					mem[out] = 0
				}
				i += 4
			case 9:
				param1 := getParam(mem, i + 1, ops[1], relativeBase)
				relativeBase += mem[param1]
				i += 2
			default:
				fmt.Println(ops)
				panic("Invalid op")
		}
	}
	return Amplifier{mem: mem, signal: signal, position: i, input: input, halt: true, relativeBase: relativeBase}
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

func inputToInt(input string) ([]int, error) {
	var out []int
	for _, val := range strings.Split(input, ",") {
		num, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		out = append(out, num)
	}
	return out, nil
}

