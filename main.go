package main

import (
	"fmt"
)

func Abs(a float64) float64 {
	if a < 0 {
		return a * -1
	}

	return a
}

func Sqrt(x float64) float64 {
	var result float64 = x
	var accuracy = 1e-8

	var prev float64 = result - ((result * result - x) / (2 * result))

	for Abs(result - prev) >= accuracy {
		prev = result
		result = result - ((result * result - x) / (2 * result))
	}

	return result
}

func main() {
	fmt.Println(Sqrt(25))
}
