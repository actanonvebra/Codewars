// Very simple, given an integer or a floating-point number, find its opposite.
package main

import "fmt"

func main() {
	fmt.Print(Opposite(5))
}

func Opposite(value int) int {
	return -value
}
