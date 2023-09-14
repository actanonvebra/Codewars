package main

import "fmt"

/*
Your car is old, it breaks easily. The shock absorbers are gone and you think it can handle about 15 more bumps before it dies totally.
Unfortunately for you, your drive is very bumpy! Given a string showing either flat road (_) or bumps (n).
If you are able to reach home safely by encountering 15 bumps or less, return Woohoo!, otherwise return Car Dead
*/

func Bump(s string) string {
	bumpsCount := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "n" {
			bumpsCount++
		}
	}
	if bumpsCount > 15 {
		return "Car Dead"
	}
	return "Wohoo!"
}

func main() {
	result := Bump("__nn_nnnn__n_n___n____nn__nnn")
	result1 := Bump("_nnnnnnn_n__n______nn__nn_nnn")
	fmt.Println(result)  //Wohoo!
	fmt.Println(result1) //Car Dead
}
