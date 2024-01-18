package main

import "fmt"

func main() {
	a := GetSize(4, 2, 6)
	b := GetSize(1, 1, 1)
	c := GetSize(1, 2, 1)
	d := GetSize(1, 2, 2)
	e := GetSize(10, 10, 10)
	fmt.Println("a'nın sonucu: ", a)
	fmt.Println("b'nin sonucu: ", b)
	fmt.Println("c'nin sonucu: ", c)
	fmt.Println("d'nin sonucu: ", d)
	fmt.Println("e'nin sonucu: ", e)
}

func GetSize(w, h, d int) [2]int {
	return [2]int{2 * (w*h + w*d + h*d), w * h * d}
}
