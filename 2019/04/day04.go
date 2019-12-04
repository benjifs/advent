package main

import (
	"fmt"
	"strconv"
)

func main() {
	passwords := generatePasswords(172930, 683082)

	fmt.Println("Number of passwords:", len(passwords))
}

func generatePasswords(start, end int) ([]int) {
	var passwords []int

	for i := start; i <= end; i++ {
		if isValid(i) {
			passwords = append(passwords, i)
		}
	}

	return passwords
}

func isValid(password int) (bool) {
	str := strconv.Itoa(password)

	if len(str) != 6 {
		return false
	}

	doubles := false
	for i := 1; i < len(str); i++ {
		prevChar := str[i - 1]
		currChar := str[i]

		if !doubles && prevChar == currChar {
			doubles = true
		}
		if currChar < prevChar {
			return false
		}
	}

	return doubles
}

