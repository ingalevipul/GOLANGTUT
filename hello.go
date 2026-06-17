// Go program to illustrate the
// concept of init() function

// Declaration of the main package
package main

// Importing package
import "fmt"

// Multiple init() function
func init() {
	fmt.Println("Welcome to init() function")
}

func init() {
	fmt.Println("Hello! init() function")
}

// Main function
func main() {
	fmt.Println("Welcome to main() function")
}

// package main
// import example.com/tut8
// func main() {
// 	const (
// 		StatusSuccess = 200
// 		StatusTimeout = 408
// 		ServerName    = "Nginx"
// 	)

// 	fmt.Println(ServerName)

// 	for {
// 		fmt.Println("vipul")
// 	}
// 	for i := 0; i <= 10; i++ {
// 		if i == 1 {
// 			break
// 		}
// 		fmt.Println(i)
// 	}
// 	i := 1

// 	for i <= 10 {
// 		fmt.Println(i)
// 		i += 1

// 	}

// 	for i := range 10 {
// 		fmt.Println(i)
// 	}

// }
