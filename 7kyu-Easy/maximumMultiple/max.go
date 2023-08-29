package main

import "fmt"

/*
Task
Given a Divisor and a Bound , Find the largest integer N , Such That ,

Conditions :
N is divisible by divisor

N is less than or equal to bound

N is greater than 0.

Notes
The parameters (divisor, bound) passed to the function are only positive values .
It's guaranteed that a divisor is Found .
*/

func MaxMultiple(d, b int) int {
	var last int = 0
	var add int = d
	for d <= b {
		last = d
		d += add
	}
	return last
}
func main() {
	result := MaxMultiple(2, 7)
	fmt.Print(result)
}
