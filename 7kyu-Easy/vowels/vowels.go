package main

import (
	"fmt"
)

/*
Trolls are attacking your comment section!

A common way to deal with this situation is to remove all of the vowels from the trolls' comments, neutralizing the threat.

Your task is to write a function that takes a string and return a new string with all vowels removed.

For example, the string "This website is for losers LOL!" would become "Ths wbst s fr lsrs LL!".

Note: for this kata y isn't considered a vowel.
*/

func DisemVowels(n string) string {
	dis := ""
	vowels := "aeiouAEIOU"
	for i := 0; i < len(n); i++ {
		notVowels := true
		for j := 0; j < len(vowels); j++ {

			if string(n[i]) == string(vowels[j]) {
				dis += ""
				notVowels = false
			}
		}
		if notVowels {
			dis += string(n[i])
		}
	}
	return dis
}

func main() {
	fmt.Print(DisemVowels("This website is for losers LOL!"))
}
