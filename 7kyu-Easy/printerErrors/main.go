package main

import (
	"fmt"
	"strconv"
)

/*

 */

func main() {

	result := PrinterErrors("aaabbbbhaijjjm")
	fmt.Println(result)
	result1 := PrinterErrors("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyzuuuuu")
	fmt.Print(result1)
}

func PrinterErrors(s string) string {
	var tmp int = 0
	var length int = len(s)
	for i := 0; i < length; i++ {
		character := s[i]
		ascii := int(character)
		if ascii < 97 || ascii > 110 {
			tmp++
		}
	}
	return strconv.Itoa(tmp) + "/" + strconv.Itoa(length)

}
