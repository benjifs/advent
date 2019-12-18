package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	lines, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	input, _ := inputToInt(lines[0])

	phases := doNPhases(input, []int{0, 1, 0, -1}, 100)
	fmt.Println("pt1:", getNDigits(phases, 8))
}

func getNDigits(in []int, n int) (string) {
	out := ""
	for i := 0; i < n; i++ {
		out += fmt.Sprintf("%d", in[i])
	}
	return out
}

func doNPhases(in, pattern []int, n int) ([]int) {
	for i := 0; i < n; i++ {
		in = doPhase(in, pattern)
	}
	return in
}

func doPhase(in, pattern []int) ([]int) {
	out := []int{}

	for len(out) != len(in) {
		p := 0
		r := 0
		for i, val := range in {
			p = ((i + 1) / (len(out) + 1)) % 4
			m := pattern[p]
			r += val * m
		}
		r = abs(r % 10)
		out = append(out, r)
	}
	return out
}

func abs(x int) (int) {
	if x < 0 {
		return x * -1
	}
	return x
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
	for _, val := range input {
		num, err := strconv.Atoi(string(val))
		if err != nil {
			return nil, err
		}
		out = append(out, num)
	}
	return out, nil
}

