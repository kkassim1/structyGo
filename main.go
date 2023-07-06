package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	// sl := []int{1, 2, 3, 55, 3, 222}
	// question 1
	// t := maxNumb(sl)
	// fmt.Println("question 1: ", t)
	//question2
	// strtest := "1k5q"
	// o := uncompress(strtest)
	// fmt.Println("question 2: ", o)
	//question3
	str := "aaabbccp"
	compressed := compress(str)
	fmt.Println("question 3: ", compressed)

	// practice
	// slic := sl[1:4]

	// fmt.Println("practie: ", slic)

}

/*
Write a function, compress, that takes in a string as an argument. The function should return a compressed version of the string where consecutive occurrences of the same characters are compressed into the number of occurrences followed by the character. Single character occurrences should not be changed.

'aaa' compresses to '3a'
'cc' compresses to '2c'
't' should remain as 't'

You can assume that the input only contains alphabetic characters.
*/

func compress(s string) string {
	result := ""
	i := 0
	j := 0

	for j < len(s) {
		if s[i] == s[j] {
			j += 1
		} else {
			count := j - i
			if count > 1 {
				result += strconv.Itoa(count) + string(s[i])
			} else {
				result += string(s[i])
			}
			i = j
		}
	}

	return result
}

/*
Write a function, uncompress, that takes in a string as an argument. The input string will be formatted into multiple groups according to the following pattern:

uncompress("2c3a1t"); // -> "ccaaat"

The function should return an uncompressed version of the string where each 'char' of a group is repeated 'number' times consecutively. You may assume that the input string is well-formed according to the previously mentioned pattern.
*/

func uncompress(s string) string {
	result := ""
	numbers := "0123456789"
	i := 0
	j := 0

	for j < len(s) {

		// if strings j is pointing to a num j++
		if strings.Contains(numbers, string(s[j])) != false {
			j += 1
		} else {
			// Atoi gets the index i and the one right before j and converts it to int
			num, _ := strconv.Atoi(s[i:j])
			// Repeat takes in the index char of located s[j] then multply by num
			result += strings.Repeat(string(s[j]), num)
			// j goes to the next num and i goes j then repat until end
			j += 1
			i = j
		}
	}
	return result
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
