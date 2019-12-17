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
	for i := 0; i < 100; i++ {
		input = append(input, 0)
	}

	game := NewGame()
	game.start(input, true)
}

type Game struct {
	min, max, pos Point

	screen map[Point]int
}

func NewGame() (Game) {
	game := Game{}
	game.min = Point{MaxInt, MaxInt}
	game.max = Point{MinInt, MinInt}
	game.screen = make(map[Point]int)
	return game
}

func (game *Game) addPoint(point Point, tile int) {
	game.screen[point] = tile
	if point.x < game.min.x {
		game.min.x = point.x
	}
	if point.x > game.max.x {
		game.max.x = point.x
	}
	if point.y < game.min.y {
		game.min.y = point.y
	}
	if point.y > game.max.y {
		game.max.y = point.y
	}
}

func (game Game) draw() {
	print("\x1b[2;0H")
	for y := game.min.y; y <= game.max.y; y++ {
		for x := game.min.x; x <= game.max.x; x++ {
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
}

func turn(dir int, clockwise bool) (int) {
	right := map[int]int{1: 4, 2: 3, 3: 1, 4: 2}
	left := map[int]int{4: 1, 3: 2, 1: 3, 2: 4}
	if clockwise {
		return right[dir]
	}
	return left[dir]
}

func (game *Game) start(input []int, render bool) {
	amp := Amplifier{mem: input, in: make(chan int), out: make(chan int), halt: make(chan bool)}
	
	go func() {
		amp.run()
		close(amp.halt)	
	}()

	in := 1
	steps := 0
	dist := make(map[Point]int)
	for i := 0; i >= 0; i++ {
		amp.in <- in
		out := <- amp.out

		look := game.pos
		switch in {
		case 1:
			look.y++
		case 2:
			look.y--
		case 3:
			look.x--
		case 4:
			look.x++
		}

		if out == 0 {
			game.addPoint(look, 1)
			in = turn(in, true)
		} else if out == 1 || out == 2 {
			game.pos = look

			if val, ok := dist[game.pos]; ok {
				steps = val
			} else {
				steps++
				dist[game.pos] = steps
			}
			if out == 1 {
				game.addPoint(game.pos, 2)
				in = turn(in, false)
			} else if out == 2 {
				game.addPoint(game.pos, 4)
				break
			}
		}
		if render {
			game.draw()
		}
	}
	fmt.Println("pt1:", dist[game.pos])
	close(amp.in)
	close(amp.out)
}

type Amplifier struct {
	mem []int
	relativeBase int

	in chan int
	out chan int
	halt chan bool
}

func getOP(code int) ([]int) {
	op := code % 100
	param1 := (code / 100) % 10
	param2 := (code / 1000) % 10
	param3 := (code / 10000) % 10

	return []int{op, param1, param2, param3}
}

func (amp Amplifier) getParam(index int, mode int) (int) {
	if mode == 0 {
		return amp.mem[index]
	}
	if mode == 2 {
		return amp.mem[index] + amp.relativeBase
	}
	return index
}

func (amp *Amplifier) run() {
	i := 0
	for {
		ops := getOP(amp.mem[i])

		op := ops[0]
		switch op {
			case 99:
				amp.halt <- true
				return
			case 1:
				param1 := amp.getParam(i + 1, ops[1])
				param2 := amp.getParam(i + 2, ops[2])
				out := amp.getParam(i + 3, ops[3])
				amp.mem[out] = amp.mem[param1] + amp.mem[param2]
				i += 4
			case 2:
				param1 := amp.getParam(i + 1, ops[1])
				param2 := amp.getParam(i + 2, ops[2])
				out := amp.getParam(i + 3, ops[3])
				amp.mem[out] = amp.mem[param1] * amp.mem[param2]
				i += 4
			case 3:
				out := amp.getParam(i + 1, ops[1])
				amp.mem[out] = <- amp.in
				i += 2
			case 4:
				out := amp.getParam(i + 1, ops[1])
				amp.out <- amp.mem[out]
				i += 2
			case 5:
				param1 := amp.getParam(i + 1, ops[1])
				param2 := amp.getParam(i + 2, ops[2])

				if amp.mem[param1] != 0 {
					i = amp.mem[param2]
				} else {
					i += 3
				}
			case 6:
				param1 := amp.getParam(i + 1, ops[1])
				param2 := amp.getParam(i + 2, ops[2])

				if amp.mem[param1] == 0 {
					i = amp.mem[param2]
				} else {
					i += 3
				}
			case 7:
				param1 := amp.getParam(i + 1, ops[1])
				param2 := amp.getParam(i + 2, ops[2])
				out := amp.getParam(i + 3, ops[3])

				if amp.mem[param1] < amp.mem[param2] {
					amp.mem[out] = 1
				} else {
					amp.mem[out] = 0
				}
				i += 4
			case 8:
				param1 := amp.getParam(i + 1, ops[1])
				param2 := amp.getParam(i + 2, ops[2])
				out := amp.getParam(i + 3, ops[3])

				if amp.mem[param1] == amp.mem[param2] {
					amp.mem[out] = 1
				} else {
					amp.mem[out] = 0
				}
				i += 4
			case 9:
				param1 := amp.getParam(i + 1, ops[1])
				amp.relativeBase += amp.mem[param1]
				i += 2
			default:
				fmt.Println(ops)
				panic("Invalid op")
		}
	}
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
