package main

import "fmt"

/*
Consider an array/list of sheep where some sheep may be missing from their place.
We need a function that counts the number of sheep present in the array (true means present).
For example,

[true,  true,  true,  false,
  true,  true,  true,  true ,
  true,  false, true,  false,
  true,  false, false, true ,
  true,  true,  true,  true ,
  false, false, true,  true]
*/

func CountSheeps(numbers []bool) int {
	var count int = 0
	for index, _ := range numbers {
		if numbers[index] {
			count++
		}
	}
	return count
}

func main() {
	sheeps := []bool{true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true}

	a := CountSheeps(sheeps)
	fmt.Print(a)
}
