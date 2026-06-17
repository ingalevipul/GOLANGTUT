package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Exercise 001:

Write a program which will find all such numbers which are divisible by 7 but
are not a multiple of 5, between 2000 and 3200 (both included). The numbers
obtained should be printed in a comma-separated sequence on a single line.

Hint: consider using strconv and strings.Join.

Tip: run `go test ./...` from this folder.
*/

// Ex001 should return all integers in [low, high] that are divisible by 7
// but not by 5, joined by commas (e.g. "112,119,126,...").
func Ex001(low, high int) string {
	var arr []string
	for i := low; i <= high; i++ {
		if i%7 == 0 && i%5 != 0 {
			arr = append(arr, strconv.Itoa(i))
		}
	}
	return strings.Join(arr, ",")

}

/*
Exercise 002:

Write a program which can compute the factorial of a given number.

  - Ex002(8) -> 40320, nil
  - Ex002(0) -> 1,     nil   (by definition)
  - Ex002(-3) -> 0,    error (negative input is not allowed)

Tip: run `go test ./...` from this folder.
*/

// Ex002 should return n! as a uint64 (or an error for negative input).
func Ex002(n int) (uint64, error) {
	var factorial int = 1

	if n < 0 {
		return 0, fmt.Errorf("error negative number")
	} else if n == 0 {
		return 1, nil

	} else {
		for n > 1 {
			factorial = factorial * n
			n -= 1
		}
		return uint64(factorial), nil
	}
}

/*
Exercise 003:

With a given integral number n, write a program to generate a map that contains
(i, i*i) such that i is an integral number between 1 and n (both included).

Example: Ex003(8) -> map[1:1 2:4 3:9 4:16 5:25 6:36 7:49 8:64]

Tip: run `go test ./...` from this folder.
*/

// Ex003 should return a map where keys are 1..n and values are key*key.
func Ex003(n int) map[int]int {
	var ret = make(map[int]int)

	for i := 1; i <= n; i++ {
		ret[i] = i * i
	}
	return ret

}

/*
Exercise 004:

Write a program which accepts a sequence of comma-separated numbers and
generates a slice out of them.

Example:
  Ex004("12, 23, 34, 45") -> []int{12, 23, 34, 45}

Tip: run `go test ./...` from this folder.
*/

// Ex004 should parse a comma-separated list of integers (whitespace is allowed
// around the numbers) and return them as []int.
func Ex004(input string) []int {
	numbers := strings.Split(input, ",")

	length := len(numbers)
	var num = make([]int, length)

	for index, v := range numbers {
		s := strings.Trim(v, " ")
		num[index], _ = strconv.Atoi(s)
	}
	return num
}
func main() {
	//fmt.Println(Ex001(2000, 3200))
	// var input int
	// _, err := fmt.Scanln(&input)
	// if err != nil {
	// 	log.Fatal("Error reading input")

	// }
	// fmt.Println(Ex002(5))
	//fmt.Println(Ex003(5))

	fmt.Println(Ex004("1,2,3,4,5"))

}
