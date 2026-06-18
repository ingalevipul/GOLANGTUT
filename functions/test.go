package main

import "fmt"

// Move it outside
func closure1() func() {
	var counter int = 0
	return func() {
		counter += 1
		fmt.Println(counter)
	}
}

func main() {
	// Call it directly
	closurefunc := closure1()
	closurefunc()

}
