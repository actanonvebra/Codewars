package main

import (
	"fmt"
	"strings"
)

func ToJadenCase(str string) string {
	var text string = ""
	for i := 0; i < len(str); i++ {
		if i == 0 {
			text += strings.ToUpper(string(str[i]))
			continue
		}
		if string(str[i]) == " " {
			text += " "
			text += strings.ToUpper(string(str[i+1]))
			i++
			continue
		}
		text += string(str[i])

	}
	return text
}

func main() {

	fmt.Print(ToJadenCase("How can 30 mirrors be real if our eyes aren't real"))
}
