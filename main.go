package main

import (
	"fmt"
)

func main() {

	sl := []int{1, 2, 3, 55, 3, 222}
	t := maxNumb(sl)

	fmt.Println(t)
}

func maxNumb(slic []int) int {
	keeptrack := -1

	for i := 0; i < len(slic); i++ {
		if slic[i] > keeptrack {
			keeptrack = slic[i]
		}
	}
	return keeptrack
}
