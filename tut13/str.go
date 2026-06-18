package main

import (
	"fmt"
	"strings"
)

func main() {
	x := "X     vipul XX"
	fmt.Println(x)
	y := strings.Trim(x, " X")
	fmt.Println(y)

}
