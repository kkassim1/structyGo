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

func maxNumb(slic []int) int {
	keeptrack := math.MinInt64

	for i := 0; i < len(slic); i++ {
		if slic[i] > keeptrack {
			keeptrack = slic[i]
		}
	}

	return keeptrack
}
