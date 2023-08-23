package main

import "fmt"

func main() {
	fmt.Print(evenOrOdd(0))
}

//Create a function that takes an integer as an argument and returns
//"Even" for even numbers or "Odd" for odd numbers.
func evenOrOdd(x int) string {
	if x%2 == 0 {
		return "Even"
	} else if x%2 != 0 {
		return "Odd"
	}
	return "0"
}
