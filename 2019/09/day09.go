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
	amp := intCode(Amplifier{mem: input, input: []int{1}})
	fmt.Println("pt1:", amp.signal)

	amp = intCode(Amplifier{mem: input, input: []int{2}})
	fmt.Println("pt2:", amp.signal)
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
