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

	input, _ := inputToInt(lines[0])
	for i := 0; i < 1000; i++ {
		input = append(input, 0)
	}
	// amp := intCode(Amplifier{mem: input, input: []int{1}})
	// fmt.Println("pt1:", amp.signal)

	// amp = intCode(Amplifier{mem: input, input: []int{2}})
	// fmt.Println("pt2:", amp.signal)

	run(input)
}

type Point struct {
	x, y int
	color int
	painted bool
}

type Robot struct {
	direction int // 0 UP, 1 RIGHT, 2, DOWN, 3 LEFT
	location Point
	points []Point
}

func run(input []int) {
	point := Point{x: 0, y: 0, color: 0, painted: false}

	robot := Robot{location: point}

	amp := Amplifier{mem: input, input: []int{}}

	colors := make(map[string]int)
	painted := make(map[string]bool)
	for {
		if amp.halt {
			break
		}
		key := strconv.Itoa(point.x) + ":" + strconv.Itoa(point.y)

		signal := 0
		if val, ok := colors[key]; ok {
			signal = val
		}
		if _, ok := painted[key]; !ok {
			painted[key] = false
		}

		amp.input = append(amp.input, signal)
		amp = intCode(amp)
		color := amp.signal
		amp = intCode(amp)
		direction := amp.signal

		if val, ok := colors[key]; ok {
			if val != color {
				painted[key] = true
			}
		}
		colors[key] = color

		if direction == 0 {
			robot.direction--
			if robot.direction < 0 {
				robot.direction = 3
			}
		} else if direction == 1 {
			robot.direction++
			if robot.direction > 3 {
				robot.direction = 0
			}
		}

		if robot.direction == 0 {
			point.y++
		} else if robot.direction == 1 {
			point.x++
		} else if robot.direction == 2 {
			point.y--
		} else if robot.direction == 3 {
			point.x--
		}
	}

	fmt.Println(len(painted))
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
}

func intCode(amp Amplifier) (Amplifier) {
	i := amp.position
	mem := amp.mem
	input := amp.input
	signal := amp.signal

	relativeBase := 0

	for i < len(mem) {
		ops := getOP(mem[i])

		op := ops[0]
		switch op {
			case 99:
				return Amplifier{mem: mem, signal: signal, position: i, input: input, halt: true}
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
				return Amplifier{mem: mem, signal: signal, position: i + 2, input: input}
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
	return Amplifier{mem: mem, signal: signal, position: i, input: input, halt: true}
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
