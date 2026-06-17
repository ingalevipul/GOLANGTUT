package main

import "fmt"

func in(nums ...int) int {
	total := 0
	for i := range nums {
		total = total + i

	}

	return total
}
func caller() func() string {
	value := func() string {
		return "Welcome! to GeeksforGeeks"
	}
	return value
}

func variadicfunc(name string, nums ...int) (string, int) {

	if nums != nil {
		total := 0
		for i := range nums {
			total += i
		}

		return name, total
	} else {
		return name, 0
	}

}

func main() {
	// nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// fmt.Println(in(nums...))
	// retfunc := caller()
	// fmt.Println(retfunc())
	fmt.Println(variadicfunc("vipul"))

}
