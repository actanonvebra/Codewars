package main

import (
	"fmt"
	"strings"
)

// Write a function which converts the input string to uppercase.
func MakeToUpper(str string) string {
	return strings.ToUpper(str)
}
func main() {
	result := MakeToUpper("HelLo, Wooorldd.")
	fmt.Print(result)
}
