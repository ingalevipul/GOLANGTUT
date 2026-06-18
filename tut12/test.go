package main

import "fmt"

func main() {
	fmt.Println("Hello")
	fmt.Printf("hello %v", "vip")

	// var num string
	// fmt.Scan(&num)
	// fmt.Printf("the number entered is %v", num)

	err := fmt.Errorf("error in program ")
	fmt.Println(err)

}
