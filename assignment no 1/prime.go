package main

import (
	"fmt"
)

func main() {
	var Input_number, Counter, Count int
	fmt.Println("Enter a number to Check for Prime:")
	fmt.Scanf("%d", &Input_number)
	if Input_number < 2 {
		println("invalid number")

	}
	for Counter = 2; Counter < Input_number/2; Counter++ {
		if Input_number%Counter == 0 {
			Count = 1
		}
	}
	if Count == 1 {
		fmt.Println("Not Prime Number")
	} else {
		fmt.Println("Prime Number")
	}

}
