package main

import (
	"fmt"
)

func MaxLength(array1, array2 []string) int {
	if len(array1) == 0 || len(array2) == 0 {
		return -1
	}
	var tmp int = 0
	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			tmp1 := len(array1[i]) - len(array2[j])
			if tmp1 < 0 {
				tmp1 = -tmp1
			}
			if tmp1 > tmp {
				tmp = tmp1
			}
		}
	}
	return tmp
}

func main() {
	a1 := []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
	a2 := []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}

	result := MaxLength(a1, a2)
	fmt.Println(result)
}
