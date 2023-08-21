package main

import "fmt"

func main() {
	fmt.Print(makeNegative(0))
}

func makeNegative(x int) int {
	if x > 0 {
		return -x
	} else if x < 0 {
		return x
	}
	return x
}
