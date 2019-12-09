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
	fmt.Println(checkLayer(img))
}

type Layer struct {
	grid [][]int
	
	count [3]int
}

type Image struct {
	layers []Layer
	width, height int
}

func checkLayer(img Image) (int) {
	res := 0
	zeroes := MaxInt
	for _, layer := range img.layers {
		for h := 0; h < len(layer.grid); h++ {
			for w := 0; w < len(layer.grid[h]); w++ {
				if layer.grid[h][w] == 0 {
					layer.count[0]++
				} else if layer.grid[h][w] == 1 {
					layer.count[1]++
				} else if layer.grid[h][w] == 2 {
					layer.count[2]++
				}
			}
		}
		if layer.count[0] < zeroes {
			zeroes = layer.count[0]
			res = layer.count[1] * layer.count[2]
		}
	}
	return res
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

