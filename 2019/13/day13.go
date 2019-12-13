package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"time"
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
	for i := 0; i < 100; i++ {
		input = append(input, 0)
	}

	game := Game{}
	game.start(input)
	fmt.Println("pt1:", game.countBlocks())

	input[0] = 2
	game.start(input)
}

type Game struct {
	w, h int
	screen map[Point]int
	ball Point
	paddle Point
	score int
}

func (game Game) countBlocks() (blocks int) {
	for _, val := range game.screen {
		if val == 2 {
			blocks++
		}
	}
	return blocks
}

func (game Game) draw() {
	print("\033[H\033[2J")
	for y := 0; y <= game.h; y++ {
		for x := 0; x <= game.w; x++ {
			char := " "
			switch game.screen[Point{x, y}] {
				case 1:
					char = "█"
				case 2:
					char = "░"
				case 3:
					char = "▀"
				case 4:
					char = "●"
			}
			fmt.Printf("%v", char)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("score: %d\n", game.score)
	time.Sleep(1 * time.Second / 100)
}

func (game *Game) start(input []int) {
	amp := Amplifier{mem: input}

	if game.screen == nil {
		game.screen = make(map[Point]int)
	}
	signal := 0
	for {
		if amp.halt {
			break
		}

		if game.paddle.x < game.ball.x {
			signal = 1
		} else if game.paddle.x > game.ball.x {
			signal = -1
		} else {
			signal = 0
		}

		amp = intCode(amp)
		x := amp.signal
		amp = intCode(amp)
		y := amp.signal
		amp = intCode(amp)
		tile := amp.signal

		if amp.requestInput {
			amp.input = append(amp.input, signal)
			game.draw()
		}

		if x == -1 && y == 0 {
			game.score = tile
		} else {
			if tile == 3 {
				game.paddle = Point{x, y}
			} else if tile == 4 {
				game.ball = Point{x, y}
			}
			game.screen[Point{x, y}] = tile
		}

		if x > game.w {
			game.w = x
		}
		if y > game.h {
			game.h = y
		}
	}
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

	requestInput bool
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
				if len(input) == 0 {
					return Amplifier{mem: mem, position: i, signal: signal, relativeBase: relativeBase, requestInput: true}
				}
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

