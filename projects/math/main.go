package main

import (
	"fmt"
	"math"
)

func t(k float64, p float64, v float64) float64 {
	mValue := m(p, v)
	return 6 / w(k, mValue) 
}

func w(k float64, m float64) float64 {
	return math.Sqrt(k / m)
}

func m(p float64, v float64) float64 {
	return p * v
}

func main() {
	nums := make([]float64, 3)
	fmt.Scan(&nums[0], &nums[1], &nums[2])
	fmt.Println(t(nums[0], nums[1], nums[2]))
}
