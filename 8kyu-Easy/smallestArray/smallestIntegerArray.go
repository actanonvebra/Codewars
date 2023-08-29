package main

import (
	"fmt"
)

func SmallestArray(n []int) int {
	var min int = n[0]
	for i := 0; i < len(n); i++ {
		if n[i] < min {
			min = n[i]
		}
	}
	return min
}

func main() {
	arr1 := []int{34, 15, 88, 2}
	arr2 := []int{34, -345, -1, 100}
	arr := []int{5, -3, -22, -7, 10, 50, -700, 40, 58, 67, 78, -777}
	result := SmallestArray(arr)
	result1 := SmallestArray(arr1)
	result2 := SmallestArray(arr2)

	fmt.Println(result)
	fmt.Println(result1) // -->2
	fmt.Print(result2)
}
