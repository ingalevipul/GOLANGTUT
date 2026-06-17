package main

import (
	"fmt"
)

func main() {

	// if age := 10; age <= 18 {
	// 	fmt.Println("u are kid")

	// } else {
	// 	fmt.Println("u are adult")

	// }

	// i := time.Now().Weekday()

	// switch i {
	// case time.Friday:
	// 	fmt.Println("monday ")

	// default:
	// 	fmt.Println("day other than monday ")

	// }

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("int datatype")

		case string:
			fmt.Println("String datatype")

		case bool:
			fmt.Println("boolean datatype")
		default:
			fmt.Println("other", t)

		}

	}

	whoAmI(true)
}
