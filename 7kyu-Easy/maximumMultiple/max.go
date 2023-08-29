package main

import "fmt"

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
