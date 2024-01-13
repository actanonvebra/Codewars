package main

import (
	"fmt"
)

/*
Finish the solution so that it sorts the passed in array of numbers.
If the function passes in an empty array or null/nil value then it should return an empty array.
*/

func main() {
	numbers := []int{4, 2, 7, 1, 9, 5}
	result := SortNumbers(numbers)
	fmt.Println(result)
}

func SortNumbers(numbers []int) []int {
	if numbers == nil {
		return []int{}
	}

	n := len(numbers)
	change := true

	for change {
		change = false
		for i := 0; i < n-1; i++ {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				change = true
			}
		}
	}
	return numbers
}
