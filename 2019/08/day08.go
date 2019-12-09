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
	fmt.Println("pt1:", getFewestZeroes(input, 25, 6))
}

func getFewestZeroes(input []int, width, height int) (int) {
	var layers [][]int

	// layers := make([][]int, height)

	// for i := range layers {
	// 	layers[i] = make([]int, width)
	// }

	// j := 0
	// for h := 0; h < height; h++ {
	// 	for w := 0; w < width; w++ {
	// 		layers[h][w] = input[j]
	// 		j++
	// 	}
	// }

	j := 0
	minLayer, minCount := 0, 99999999
	zeroes := 0
	layers = append(layers, []int{})
	for i, vals := range input {
		if vals == 0 {
			zeroes++
		}
		layers[j] = append(layers[j], vals)

		fmt.Println(i)

		if (i + 1) % (width * height) == 0 {
			layers = append(layers, []int{})
			if zeroes < minCount {
				minLayer = j
				minCount = zeroes
			}
			fmt.Println(j, zeroes)
			j++
			zeroes = 0
		}
	}


	countOne := 0
	countTwo := 0
	for _, pixel := range layers[minLayer] {
		if pixel == 1 {
			countOne++
		}
		if pixel == 2 {
			countTwo++
		}
	}

	return countOne * countTwo

	// minLayer, minCount := 0, 99999
	// for i, layer := range layers {
	// 	zeroes := 0
	// 	for _, pixel := range layer {
	// 		if pixel == 0 {
	// 			zeroes++
	// 		}
	// 	}

	// 	if zeroes < minCount {
	// 		minLayer = i
	// 		minCount = zeroes
	// 	}
	// }

	// fmt.Println("minlayer", minLayer, minCount)

	// countOne := 0
	// countTwo := 0
	// for _, pixel := range layers[minLayer] {
	// 	if pixel == 1 {
	// 		countOne++
	// 	}
	// 	if pixel == 2 {
	// 		countTwo++
	// 	}
	// }

	// fmt.Println(layers)

	// return countOne * countTwo
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

