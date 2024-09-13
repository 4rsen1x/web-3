package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input int
	fmt.Scan(&input)

	str := strconv.Itoa(input)
	result := ""

	for _, char := range str {
		num, _ := strconv.Atoi(string(char))
		square := num * num
		result += strconv.Itoa(square)
	}

	fmt.Println(result)
}