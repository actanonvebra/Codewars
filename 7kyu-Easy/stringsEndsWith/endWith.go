package main

import "fmt"

/*
Complete the solution so that it returns true if the first argument(string)
passed in ends with the 2nd argument (also a string).

Examples:

solution("abc", "bc") // returns true
solution("abc", "d") // returns false

*/

func Solution(str, ending string) bool {
	if len(ending) > len(str) {
		return false
	}
	isTrue := false
	count := 0
	index := len(str) - 1 - len(ending)
	reverse := ""
	for count < len(ending) {
		reverse += string(str[index+1])
		index++
		count++
	}
	if reverse == ending {
		isTrue = true
	}
	return isTrue
}

func main() {
	result := Solution("", "t")
	fmt.Println(result) // false

}
