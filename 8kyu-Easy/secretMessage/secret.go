package main

import "fmt"

func Greet(name string) string {
	if name == "Johnny" {
		return "Hello, my love!"
	}
	return "Hello," + " " + name + "!"
}

func main() {
	var firstName = "Johnny"
	var secondName = "Alfredo"

	fmt.Println(Greet(firstName)) //Hello, my love!
	fmt.Print(Greet(secondName))  //Hello, Alfredo!
}
