package main

import "fmt"

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
