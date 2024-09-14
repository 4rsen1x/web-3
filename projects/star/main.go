package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)
	for i, char := range input {
		fmt.Print(string(char))
		if i < len(input) - 1 {
			fmt.Print("*")
		}
	}
}
