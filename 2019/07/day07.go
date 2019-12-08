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

	best := -1
	input, _ := inputToInt(lines[0])
	seq := []int{0, 1, 2, 3, 4}
	for _, phaseSettings := range getPhaseSettings(seq) {
		out := runWithPhaseSetting(input, phaseSettings)
		if best < out {
			best = out
		}
	}
	fmt.Println("pt1:", best)

	best = -1
	seq2 := []int{5, 6, 7, 8, 9}
	for _, phaseSettings := range getPhaseSettings(seq2) {
		out := runWithPhaseSetting2(input, phaseSettings)
		if best < out {
			best = out
		}
	}
	fmt.Println("pt2:", best)
}

func runWithPhaseSetting2(input, phaseSettings []int) (int) {
	signal := 0

	var amps []Amplifier
	for _, phaseSetting := range phaseSettings {
		amps = append(amps, Amplifier{mem: input, input: []int{phaseSetting}})
	}

	i := 0
	for {
		amps[i].input = append(amps[i].input, signal)
		amps[i] = intCode(amps[i])
		signal = amps[i].signal

		if i == 4 && amps[i].halt {
			break
		}
		i = (i + 1) % len(amps)
	}
	return signal
}

func runWithPhaseSetting(input, phaseSettings []int) (int) {
	signal := 0
	for _, phaseSetting := range phaseSettings {
		amp := intCode(Amplifier{mem: input, input: []int{phaseSetting, signal}})
		signal = amp.signal
	}
	return signal
}

func getPhaseSettings(seq []int) (result [][]int) {
	if len(seq) == 1 {
		return append(result, seq)
	}
	for i, char := range seq {
		var tmp []int
		tmp = append(tmp, seq[:i]...)
		tmp = append(tmp, seq[i+1:]...)

		for _, p := range getPhaseSettings(tmp) {
			result = append(result, append(p, char))
		}
	}
	return result
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

type Amplifier struct {
	mem []int
	position int
	signal int
	input []int
	halt bool
}

func intCode(amp Amplifier) (Amplifier) {
	var diagnostic int

	i := amp.position
	mem := amp.mem
	input := amp.input

	for i < len(mem) {
		ops := getOP(mem[i])

		op := ops[0]
		switch op {
			case 99:
				return Amplifier{mem: mem, signal: diagnostic, position: i, input: amp.input, halt: true}
			case 1:
				param1 := getParam(mem, i + 1, ops[1])
				param2 := getParam(mem, i + 2, ops[2])
				out := getParam(mem, i + 3, 1)
				mem[out] = param1 + param2
				i += 4
			case 2:
				param1 := getParam(mem, i + 1, ops[1])
				param2 := getParam(mem, i + 2, ops[2])
				out := getParam(mem, i + 3, 1)
				mem[out] = param1 * param2
				i += 4
			case 3:
				if len(input) == 0 {
					return Amplifier{mem: mem, signal: diagnostic, position: i, input: input}
				}
				out := getParam(mem, i + 1, 1)
				mem[out], input = input[0], input[1:]
				i += 2
			case 4:
				out := getParam(mem, i + 1, ops[1])
				diagnostic = out
				i += 2
			case 5:
				param1 := getParam(mem, i + 1, ops[1])
				param2 := getParam(mem, i + 2, ops[2])
				if param1 != 0 {
					i = param2
				} else {
					i += 3
				}
			case 6:
				param1 := getParam(mem, i + 1, ops[1])
				param2 := getParam(mem, i + 2, ops[2])
				if param1 == 0 {
					i = param2
				} else {
					i += 3
				}
			case 7:
				param1 := getParam(mem, i + 1, ops[1])
				param2 := getParam(mem, i + 2, ops[2])
				out := getParam(mem, i + 3, 1)
				if param1 < param2 {
					mem[out] = 1
				} else {
					mem[out] = 0
				}
				i += 4
			case 8:
				param1 := getParam(mem, i + 1, ops[1])
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
	return Amplifier{mem: mem, signal: diagnostic, position: i, input: amp.input, halt: true}
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

