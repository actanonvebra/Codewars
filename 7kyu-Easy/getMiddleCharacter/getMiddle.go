package main

import "fmt"

/*
DESCRIPTION:
You are going to be given a word. Your job is to return the middle character of the word. If the word's length is odd, return the middle character. If the word's length is even, return the middle 2 characters.

#Examples:

Kata.getMiddle("test") should return "es"

Kata.getMiddle("testing") should return "t"

Kata.getMiddle("middle") should return "dd"

Kata.getMiddle("A") should return "A"
*/

func GetMiddle(s string) string {
	var size int = len(s)
	if size%2 == 0 {
		return s[size/2-1 : size/2+1]
	} else if size%2 == 1 {
		return s[size/2 : size/2+1]
	}
	return s
}

func main() {
	fmt.Println(GetMiddle("A"))       // A
	fmt.Println(GetMiddle("ABBA"))    // BB
	fmt.Println(GetMiddle("ibrahim")) // a
}
