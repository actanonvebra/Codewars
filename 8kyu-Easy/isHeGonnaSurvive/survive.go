package main

import "fmt"

func Hero(bullets, dragons int) bool {
	if dragons*2 <= bullets {
		return true
	}
	return false
}

func main() {
	var result1 = Hero(10, 2)
	var result2 = Hero(5, 4)
	fmt.Print(result1) // true
	fmt.Print(result2) // false
}
