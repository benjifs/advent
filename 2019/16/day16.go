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
	fmt.Println("pt1:", sliceToString(phases[:8]))
	fmt.Println("pt2:", getSignal(input))
}

func getSignal(in []int) (string) {
	offset, _ := strconv.Atoi(sliceToString(in[:7]))

	input := make([]int, len(in) * 10000)
	for i := 0; i < 10000; i++ {
		for j, char := range in {
			input[j + len(in) * i] = char
		}
	}

	signal := input[offset:len(input)]
	for i := 0; i < 100; i++ {
		sum := 0
		for j := len(signal) - 1; j >= 0; j-- {
			sum += signal[j]
			signal[j] = abs(sum % 10)
		}
	}

	return sliceToString(signal[:8])
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

func sliceToString(slice []int) (string) {
	out := ""
	for _, n := range slice {
		out += fmt.Sprintf("%d", n)
	}
	return out
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

