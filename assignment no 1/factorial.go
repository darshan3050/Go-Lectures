package main

import "fmt"

func main() {

	var Input_number, Output int
	Output = 1

	fmt.Print("Enter Number = ")
	fmt.Scanln(&Input_number)

	for i := 1; i <= Input_number; i++ {
		Output = Output * i
	}
	fmt.Printf("The Factorial is   = %d", Output)
}
