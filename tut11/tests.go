package main

import "fmt"

type person struct {
	name string
	age  int
}

func change(p person) {
	p.name = "Kalki"
	p.age = 100
}

func main() {
	p1 := person{"Vipul", 21}
	change(p1)
	fmt.Println(p1.name)
	fmt.Println(p1.age)
}
