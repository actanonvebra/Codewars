package main

import (
	"fmt"
)

func main() {
	var name string = "j  oh n    "
	fmt.Print(ReplaceSpaces(name)) // -->john
}

/*
Write a function that removes the spaces from the string
then return the resultant string.

Examples:

Input -> Output
"8 j 8   mBliB8g  imjB8B8  jl  B" -> "8j8mBliB8gimjB8B8jlB"
"8 8 Bi fk8h B 8 BB8B B B  B888 c hl8 BhB fd" -> "88Bifk8hB8BB8BBBB888chl8BhBfd"
"8aaaaa dddd r     " -> "8aaaaaddddr
*/

func ReplaceSpaces(word string) string {
	var temp string = ""
	for index := 0; index < len(word); index++ {
		if string(word[index]) != " " {
			temp += string(word[index])
		}
		continue
	}
	return temp
}
