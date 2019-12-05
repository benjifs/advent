package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "errors"
	"strconv"
)

func main() {
	lines, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		input, _ := inputToInt(line)

		intCode(input, 1)
	}
}

func getOP(o int) ([]int) {
	var out []int

	if o > 99 {
		out = append(out, o % 100)
		o = o / 100
		for o > 9 {
			out = append(out, o % 10)
			o = o / 10
		}
		if o > 0 {
			out = append(out, o)
		}
	} else {
		out = append(out, o)
	}

	for i := len(out); i < 4; i++ {
		out = append(out, 0)
	}

	return out
}

func intCode(input []int, in int) ([]int, error) {
	mem := make([]int, len(input))
	copy(mem, input)

	i := 0
	for i < len(mem) {
		ops := getOP(mem[i])
		op := ops[0]

		if op == 99 {
			break
		}

		if op == 1 || op == 2 {
			a := mem[i + 1]
			b := mem[i + 2]
			c := mem[i + 3]

			var param1, param2 int

			if ops[1] == 0 {
				param1 = mem[a]
			} else {
				param1 = a
			}
			if ops[2] == 0 {
				param2 = mem[b]
			} else {
				param2 = b
			}

			if op == 1 {
				mem[c] = param1 + param2
			} else {
				mem[c] = param1 * param2
			}

			i += 4
		} else if op == 3 || op == 4 {
			a := mem[i + 1]

			if op == 3 {
				mem[a] = in
			} else {
				fmt.Println(mem[a])
			}

			i += 2
		} else {
			panic("Invalid op")
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

