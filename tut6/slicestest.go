package main

import "fmt"

// func Pic(dx, dy int) [][]uint8 {

// }

func main() {
	// var matrix [][]int // Nil slice with 0 length
	// dx := 3
	// dy := 3

	// for i := 1; i <= dx; i++ {

	// 	var matrix2 = make([]int, 0)

	// 	for j := 1; j <= dy; j++ {
	// 		matrix2 = append(matrix2, i*j)
	// 	}

	// 	matrix = append(matrix, matrix2)

	// }

	// fmt.Println(matrix)

	arr1 := []int{1, 2, 3}
	for _, i := range arr1 {
		fmt.Println(i)
	}

	for i, r := range "Hello" {
		fmt.Printf("%d %c\n", i, r)
	}
}
