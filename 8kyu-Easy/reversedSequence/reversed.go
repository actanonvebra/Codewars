package main

import "fmt"

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
