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
	game.start(input, false)
	fmt.Println("pt1:", game.countBlocks())

	input[0] = 2
	game.start(input, true)
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
	time.Sleep(1 * time.Second / 50)
}

func (game *Game) start(input []int, render bool) {
	amp := Amplifier{mem: input, in: make(chan int), out: make(chan int), halt: make(chan bool)}

	if game.screen == nil {
		game.screen = make(map[Point]int)
	}
	
	go func() {
		amp.run()
		close(amp.halt)	
	}()

	signal := 0
	running := true
	for running {
		select {
			case <- amp.halt:
				running = false
			case amp.in <- signal:
				if render {
					game.draw()
				}
			case x := <- amp.out:
				y := <- amp.out
				tile := <- amp.out

				if x == -1 && y == 0 {
					game.score = tile
				} else {
					if tile == 3 {
						game.paddle = Point{x, y}
					} else if tile == 4 {
						game.ball = Point{x, y}
						if game.paddle.x < game.ball.x {
							signal = 1
						} else if game.paddle.x > game.ball.x {
							signal = -1
						} else {
							signal = 0
						}
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
	if render {
		game.draw()
	}
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

