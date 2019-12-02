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
		out, _ := intCode(input)
		fmt.Println(out[0])
	}
}

func intCode(input []int) ([]int, error) {
	for i := 0; i < len(input); i += 4 {
		op := input[i]

		if op == 99 {
			break
		}

		a := input[i + 1]
		b := input[i + 2]
		c := input[i + 3]
		if op == 1 {
			input[c] = input[a] + input[b]
		} else if op == 2 {
			input[c] = input[a] * input[b]
		} else {
			return nil, errors.New("Invalid op")
		}
	}
	return input, nil
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
	vals := strings.Split(input, ",")

	var out []int
	for _, val := range vals {
		num, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		out = append(out, num)
	}
	return out, nil
}

