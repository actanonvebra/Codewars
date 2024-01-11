package main

import "fmt"

func main() {
	first := Add(20) //fonksiyonu temsil ediyor.
	second := first(10)
	fmt.Print(second)
}

func Add(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}
