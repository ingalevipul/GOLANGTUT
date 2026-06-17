package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c int

	fmt.Println(a, b)

	c = a
	a = b
	b = c

	fmt.Println(a, b)
}
