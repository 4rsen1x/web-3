package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)

	max := '0'

	for _, char := range input {
		if char > max {
			max = char
		}
	}

	fmt.Printf(string(max))
}