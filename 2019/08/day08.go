package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MaxInt = int(^uint(0) >> 1)

func main() {
	lines, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	input, _ := inputToInt(lines[0])

	img := generateImage(input, 25, 6)
	layer := getFewestZeroes(img)
	fmt.Println(checkLayer(img.layers[layer]))


}

type Layer struct {
	grid [][]int
}

type Image struct {
	layers []Layer
	width, height int
}

func checkLayer(layer Layer) (int) {
	countOne, countTwo := 0, 0
	for h := 0; h < len(layer.grid); h++ {
		for w := 0; w < len(layer.grid[h]); w++ {
			if layer.grid[h][w] == 1 {
				countOne++
			} else if layer.grid[h][w] == 2 {
				countTwo++
			}
		}
	}
	return countOne * countTwo
}

func getFewestZeroes(img Image) (int) {
	zeroLayer := 0
	zeroes := MaxInt
	for l, layer := range img.layers {
		count := 0
		for h := 0; h < len(layer.grid); h++ {
			for w := 0; w < len(layer.grid[h]); w++ {
				if layer.grid[h][w] == 0 {
					count++
				}
			}
		}
		if count < zeroes {
			zeroLayer = l
			zeroes = count
		}
	}
	return zeroLayer
}

func generateLayer(pixels []int, width, height int) (Layer) {
	layer := Layer{}

	for h := 0; h < height; h++ {
		layer.grid = append(layer.grid, []int{})
		for w := 0; w < width; w++ {
			layer.grid[h] = append(layer.grid[h], pixels[h * width + w])
		}
	}
	return layer
}

func generateImage(pixels []int, width, height int) (Image) {
	img := Image{width: width, height: height}

	for len(pixels) > 0 {
		layer := pixels[:(width * height)]
		pixels = pixels[(width * height):]

		img.layers = append(img.layers, generateLayer(layer, width, height))
	}

	return img
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

