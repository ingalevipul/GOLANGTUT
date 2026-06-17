package main

import (
	"fmt"
	"sort"
)

func printpointers(ptrs ...*[5]int) {
	for _, v := range ptrs {
		fmt.Println(v)
	}
}
func increptrarr(ptr *[5]int) {
	elm := ptr
	elm[0] = 100
	fmt.Println(elm)

}
func slices() {
	intarr := [5]int{12, 34, 55, 66, 43}

	slice := intarr[:]
	//intarr[1] = 1000

	slice = append(slice, 1, 2, 4, 5, 5, 6)
	// fmt.Println(intarr, slice)

	// fmt.Printf("address of slice %p add of Arr %p\n", &slice[0], &intarr)
	sort.Ints(slice)
	fmt.Println(slice)
	fmt.Printf("address of slice %p add of Arr %p\n", &slice[0], &intarr)

}
func main() {
	slices()
	// var source = [5]int{10, 20, 30, 40, 50}
	// arr1 := &source
	// fmt.Println(*arr1)
	// increptrarr(arr1)
	// fmt.Println(*arr1)
	// // Copying by direct assignment
	// var destination [5]int = source

	// printpointers(&source, &destination)
	// source[0] = 100
	// fmt.Println(source, destination)
	// var colors [4]string

	// fmt.Println(colors)

	// arr1 := [3][3]int{{1, 2, s3}, {1, 3, 6}}
	// fmt.Println(arr1)

	// var arr []int
	// var arr1 = make([]int, 0, 4)
	// for i := 1; i <= 5; i++ {
	// 	arr1 = append(arr1, i)
	// }
	// arr1 = append(arr1, 6)

	// fmt.Println(arr1, len(arr1))

	// s := make([]int, 0, 5)

	// for i := 1; i <= 6; i++ {
	// 	fmt.Printf("before append: len=%d cap=%d\n", len(s), cap(s))
	// 	s = append(s, i)
	// }

	// arr := [5]int{10, 20, 30, 40, 50}

	// s1 := arr[1:4] // [20 30 40]
	// s2 := arr[2:5] // [30 40 50]

	// s1[1] = 999
	// fmt.Println(arr) // [10 20 999 40 50]
	// fmt.Println(s1)  // [20 999 40]
	// fmt.Println(s2)  // [999 40 50]

	// var arr = make([]int, 0, 5)
	// var arr1 = make([]int, 5)
	// var arr2 = []int{1, 3, 4, 5}

	// arr = append(arr, 1, 2, 3, 4, 5)
	// copy(arr1, arr)

	// arr[1] = 1000

	// fmt.Println(arr, arr1)
	// fmt.Println(arr2)

}
