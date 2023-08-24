package main

import "fmt"

func main() {
	var nums = [4]int{1, -4, 7, 12}
	fmt.Print(SumOfPositive(nums[:])) // -->20
}

/*
You get an array of numbers, return the sum of all of the positives ones.

Example [1,-4,7,12] => 1 + 7 + 12 = 20

Note: if there is nothing to sum, the sum is default to 0.
*/
func SumOfPositive(arr []int) int {
	var sum int = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < -arr[i] {
			continue
		}
		sum += arr[i]
	}
	return sum
}
