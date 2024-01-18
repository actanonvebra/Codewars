package main

import "fmt"

/*
DESCRIPTION:
Create a combat function that takes the player's current health and the amount of damage recieved
and returns the player's new health.
Health can't be less than 0.
*/

func main() {
	//test
	a := combat(100, 102)
	fmt.Print(a)
}

func combat(health, damage float64) float64 {
	var result float64 = health - damage
	if result < 0 {
		return 0
	}
	return result
}
