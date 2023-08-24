package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(Solution("scoby", "doo"))
}

/*
Given 2 strings, a and b, return a string of the form short+long+short, with the shorter string on the outside
and the longer string on the inside.
The strings will not be the same length, but they may be empty ( zero length ).
Hint for R users: The length of string is not always the same as the number of characters
*/

func Solution(a, b string) string {
	aReplace := strings.Replace(a, " ", "", -1)
	bReplace := strings.Replace(b, " ", "", -1)
	if len(aReplace) > len(bReplace) {
		return bReplace + aReplace + bReplace
	}
	return aReplace + bReplace + aReplace
}
