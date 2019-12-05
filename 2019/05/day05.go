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
	for _, line := range lines {
		input, _ := inputToInt(line)
		
		_, diagnostic := intCode(input, 1)
		fmt.Println("pt1 diagnostic:", diagnostic)

		_, diagnostic = intCode(input, 5)
		fmt.Println("pt2 diagnostic:", diagnostic)
	}
}

func getOP(code int) ([]int) {
	op := code % 100
	param1 := (code / 100) % 10
	param2 := (code / 1000) % 10
	param3 := (code / 10000) % 10

	return []int{op, param1, param2, param3}
}

func getParam(mem []int, index int, mode int) (int) {
	if mode == 0 {
		return mem[mem[index]]
	}
	return mem[index]
}

func intCode(instructions []int, input int) ([]int, int) {
	mem := make([]int, len(instructions))
	copy(mem, instructions)

	var diagnostic int

	i := 0
	for i < len(mem) {
		ops := getOP(mem[i])

		op := ops[0]
		if op == 99 {
			break
		}

		param1 := getParam(mem, i + 1, ops[1])
		switch op {
			case 1:
				param2 := getParam(mem, i + 2, ops[2])
				out := getParam(mem, i + 3, 1)
				mem[out] = param1 + param2
				i += 4
			case 2:
				param2 := getParam(mem, i + 2, ops[2])
				out := getParam(mem, i + 3, 1)
				mem[out] = param1 * param2
				i += 4
			case 3:
				out := getParam(mem, i + 1, 1)
				mem[out] = input
				i += 2
			case 4:
				diagnostic = mem[mem[i + 1]]
				i += 2
			case 5:
				param2 := getParam(mem, i + 2, ops[2])
				if param1 != 0 {
					i = param2
				} else {
					i += 3
				}
			case 6:
				param2 := getParam(mem, i + 2, ops[2])
				if param1 == 0 {
					i = param2
				} else {
					i += 3
				}
			case 7:
				param2 := getParam(mem, i + 2, ops[2])
				out := getParam(mem, i + 3, 1)
				if param1 < param2 {
					mem[out] = 1
				} else {
					mem[out] = 0
				}
				i += 4
			case 8:
				param2 := getParam(mem, i + 2, ops[2])
				out := getParam(mem, i + 3, 1)
				if param1 == param2 {
					mem[out] = 1
				} else {
					mem[out] = 0
				}
				i += 4
			default:
				fmt.Println(ops)
				panic("Invalid op")
		}
	}
	return mem, diagnostic
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
