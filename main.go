package main

import (
	"fmt"
	"math"
)

func main() {
	sl := []int{1, 2, 3, 55, 3, 222}
	t := maxNumb(sl)
	fmt.Println(t)
}

/*
Write a function, maxValue, that takes in a vector of numbers as an argument. The function should return the largest number in the vector.
Solve this without using any built-in methods.
You can assume that the vector is non-empty.
*/
func maxNumb(slic []int) int {
	keeptrack := math.MinInt64

	for i := 0; i < len(slic); i++ {
		if slic[i] > keeptrack {
			keeptrack = slic[i]
		}
	}

	return keeptrack
}
