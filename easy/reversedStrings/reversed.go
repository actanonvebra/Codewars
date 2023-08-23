package main

import "fmt"

func main() {
	var name string = "serhat"
	fmt.Print(Solution(name))
}

//Complete the solution so that it reverses the string passed into it.
// 'world'  =>  'dlrow'
// 'word'   =>  'drow'
func Solution(word string) string {
	var reverseWord string = ""
	var count int = 0
	for i := 0; i < len(word); i++ {
		count++
		reverseWord += string(word[len(word)-count])
	}
	return reverseWord
}
