package main

import "fmt"

func main() {
	fmt.Println(RemoveChar("scoby doo"))
}

/*
It's pretty straightforward.
Your goal is to create a function that removes the first and last characters of a string.
You're given one parameter, the original string.
You don't have to worry with strings with less than two characters.
*/
func RemoveChar(word string) string {
	length := len(word)
	return word[1 : length-1]
}
