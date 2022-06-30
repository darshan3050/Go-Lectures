package main

import "fmt"

func main() {
	var a int
	fmt.Println("select from 1 to 4")
	fmt.Scanf("%d", &a)
	switch {
	case a == 1:
		fmt.Println("first")
	case a == 2:
		fmt.Println("second")
	case a == 3:
		fmt.Println("third")
	case a == 4:
		fmt.Println("fourth")
	default:
		fmt.Println("invalid")

	}
}
