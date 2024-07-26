package main

import (
	"fmt"
)

func dd(a map[string]int) {
	a["fddf"] = 3334
}

func main() {
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)
}
