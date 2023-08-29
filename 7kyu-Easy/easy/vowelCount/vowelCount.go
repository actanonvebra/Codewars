package main

import "fmt"

/*
Return the number (count) of vowels in the given string.

We will consider a, e, i, o, u as vowels for this Kata (but not y).

The input string will only consist of lower case letters and/or spaces.
*/
func GetCount(str string) (count int) {
	count = 0
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(vowels); j++ {
			if str[i] == byte(vowels[j]) {
				count++
			}
		}
	}
	return count
}
func main() {
	fmt.Print(GetCount("aabcd")) // -->2
}
