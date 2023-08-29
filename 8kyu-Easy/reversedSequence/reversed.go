package main

import "fmt"

/*
Build a function that returns an array of integers from n to 1 where n>0.

Example : n=5 --> [5,4,3,2,1]
*/
func ReverseSeq(n int) []int {
	arr := make([]int, n)
	count := 0
	for n >= 1 {
		arr[count] = n
		count++
		n--
	}
	return arr
}

func main() {
	result := ReverseSeq(5)
	fmt.Print(result)
}
