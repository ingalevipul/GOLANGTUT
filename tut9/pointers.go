package main

import "fmt"

func update(num *int) {
	*num = 5

}
func main() {
	num := 1
	ptr1 := &num
	ptr2 := &ptr1
	fmt.Println(num)
	fmt.Println(**ptr2)

	update(&num)
	fmt.Println(num)
}
