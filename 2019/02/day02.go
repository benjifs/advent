package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"errors"
	"strconv"
)

func main() {
	lines, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		input, err := inputToInt(line)
		if err != nil {
			panic(err)
		}

		// PART 01
		input[1] = 12
		input[2] = 2
		out, _ := intCode(input)
		fmt.Printf("part 1: %d\n", out[0])

		// PART 02
		for noun := 0; noun <= 99; noun++ {
			for verb := 0; verb <= 99; verb++ {
				input[1] = noun
				input[2] = verb
				out, _ := intCode(input)

				if out[0] == 19690720 {
					fmt.Printf("part 2: %d\n", (100 * noun) + verb)
					return
				}
			}
		}
	}
}

func intCode(input []int) ([]int, error) {
	mem := make([]int, len(input))
	copy(mem, input)
	for i := 0; i < len(mem); i += 4 {
		op := mem[i]

		if op == 99 {
			break
		}

		a := mem[i + 1]
		b := mem[i + 2]
		c := mem[i + 3]
		if op == 1 {
			mem[c] = mem[a] + mem[b]
		} else if op == 2 {
			mem[c] = mem[a] * mem[b]
		} else {
			return nil, errors.New("Invalid op")
		}
	}
	return mem, nil
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

