package main

import "fmt"

func main() {
	fmt.Print(BoolToWord(true))  // --> Yes
	fmt.Print(BoolToWord(false)) // --> No

}

//Complete the method that takes a boolean value and return a "Yes" string for true, or a "No" string for false.
func BoolToWord(word bool) string {
	if word {
		return "Yes"
	}
	return "No"
}
