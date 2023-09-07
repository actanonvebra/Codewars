package main

import (
	"fmt"
	"strings"
)

/*
This time no story, no theory. The examples below show you how to write function accum:

Examples:
accum("abcd") -> "A-Bb-Ccc-Dddd"
accum("RqaEzty") -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
accum("cwAt") -> "C-Ww-Aaa-Tttt"
The parameter of accum is a string which includes only letters from a..z and A..Z.
*/

func Accum(s string) string {
	var slice []string
	for i := 0; i < len(s); i++ {
		slice = append(slice, string(s[i]))
	}
	var r string = ""
	for i := 0; i < len(slice); i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				r += strings.ToUpper(slice[i])
				continue
			}
			r += strings.ToLower(slice[i])
		}
		if i+1 == len(slice) {
			break
		}
		r += "-"
	}
	return r
}

func main() {
	fmt.Print(Accum("Zyp"))
}
