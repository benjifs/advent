package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lines, err := readInput("input.txt")
	// lines, err := readInput("test_input.txt")
	if err != nil {
		panic(err)
	}

	planets := make(map[string][]string)
	for _, line := range lines {
		values := strings.Split(line, ")")

		var obj []string
		if val, ok := planets[values[1]]; ok {
			obj = val
		}

		obj = append(obj, values[0])
		planets[values[1]] = obj
	}


	i := 0
	for k, _ := range planets {
		tmp := maptree(k, planets)
		fmt.Println(k, tmp)

		i+= tmp
	}
	fmt.Println(i)
}

func maptree(name string, planets map[string][]string) (int) {
	if len(planets[name]) == 0 {
		return 0
	}
	for _, o := range planets[name] {
		return maptree(o, planets) + 1
	}
	return 0
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

