package main

import (
	"fmt"
)

func main() {
	var n, i, c int
	fmt.Scanf("%d", &n)

	if n < 2 {
		println("invalid number")

	}
	for i = 2; i < n/2; i++ {
		if n%i == 0 {
			c = 1
		}
	}
	if c == 1 {
		fmt.Println("Not Prime Number")
	} else {
		fmt.Println("Prime Number")
	}

}
