package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	fmt.Printf("%d\t %d\n", a, b)
	if a > b {
		fmt.Printf("%d is greater than %d", a, b)

	} else {
		fmt.Printf("%d is greater than %d", b, a)
	}

}
