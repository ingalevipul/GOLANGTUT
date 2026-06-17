package main

import (
	"fmt"
)

func main() {
	var duct = make(map[string]string)
	duct["name"] = "vipul"
	duct["clg"] = "tsec"
	duct["branch"] = "AI&DS"

	// fmt.Println(duct)

	// delete(duct, "name")
	// fmt.Println(duct)
	// clear(duct)
	// fmt.Println(duct)

	// v, ok := duct["name"]
	// if ok {
	// 	fmt.Println(ok, v)

	// } else {
	// 	fmt.Println(v)
	// }

	// for k, v := range duct {
	// 	fmt.Println(k, v)
	// }
	val, ok := duct["a"]
	if !ok {
		fmt.Println(val, ok == true)
	}

}
