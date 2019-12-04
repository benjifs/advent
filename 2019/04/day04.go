package main

import (
	"fmt"
	"strconv"
)

func main() {
	passwords := generatePasswords(172930, 683082, false)
	fmt.Println("Number of passwords (pt1):", len(passwords))

	passwords = generatePasswords(172930, 683082, true)
	fmt.Println("Number of passwords (pt2):", len(passwords))
}

func generatePasswords(start, end int, checkDoubles bool) ([]int) {
	var passwords []int

	for i := start; i <= end; i++ {
		if isValid(i) {
			if !checkDoubles {
				passwords = append(passwords, i)
			} else if hasDoubles(i) {
				passwords = append(passwords, i)
			}
		}
	}

	return passwords
}

func hasDoubles(password int) (bool) {
	str := strconv.Itoa(password)

	var split []string
	part := string(str[0])
	for i := 1; i < len(str); i++ {
		if str[i - 1] == str[i] {
			part += string(str[i])
		} else {
			split = append(split, part)
			part = string(str[i])
		}
	}
	split = append(split, part)

	for _, s := range split {
		if len(s) == 2 {
			return true
		}
	}

	return false
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

