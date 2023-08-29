package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4}
	fmt.Print(SquareSum(array))
}

/*
Complete the square sum function so that it squares
each number passed into it and then sums the results together.

For example, for [1, 2, 2] it should return 9 because
1^2+2^2+2^2=9
*/
func SquareSum(numbers []int) int {
	var sum int = 0
	for index := range numbers {
		sum += numbers[index] * numbers[index]
	}
	return sum
}
