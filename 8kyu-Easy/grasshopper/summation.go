package main

import "fmt"

func main() {
	fmt.Println(Summation0(2))
	fmt.Println(Summation1(2))
}

/*
Write a program that finds the summation of every number from 1 to num.
The number will always be a positive integer greater than 0.

For example (Input -> Output):
2 -> 3 (1 + 2)
8 -> 36 (1 + 2 + 3 + 4 + 5 + 6 + 7 + 8)

*/

func Summation0(n int) int {
	var sum int = 0
	for i := 1; i <= n; {
		sum += i
		i++
	}
	return sum
}

//Daha hoş olanı.
func Summation1(n int) int {
	return n * (n + 1) / 2
}
