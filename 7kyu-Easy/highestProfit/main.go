package main

import (
	"fmt"
)

func main() {

	/*Story
	Ben has a very simple idea to make some profit: he buys something and sells it again.
	Of course, this wouldn't give him any profit at all
	if he was simply to buy and sell it at the same price.
	Instead, he's going to buy it for the lowest possible price and sell it at the highest.

	Task
	Write a function that returns both the minimum and maximum number of the given list/array.

	Examples (Input --> Output)
	[1,2,3,4,5] --> [1,5]
	[2334454,5] --> [5,2334454]
	[1]         --> [1,1]
	Remarks
	All arrays or lists will always have at least one element
	so you don't need to check the length.
	Also, your function will always get an array or a list, you don't have to check for null
	undefined or similar.
	*/

	// EASY
	// array := []int{5, 6, 7, 3, 1, 2}
	// sort.Ints(array)
	// fmt.Println(array)
	//-------------------------------------------

	arr := []int{2424234, 6, 7, 3, 1, 2}
	result := MinMax(arr)
	fmt.Print(result)
}
func MinMax(arr []int) [2]int {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	var first int = arr[0]
	var last int = arr[len(arr)-1]
	return [2]int{first, last}
}
