package main

import "fmt"

func main() {
	var First_integer = 0
	var Second_integer = 1
	var sum, Counter, Input_number, Temperory_Value int
	fmt.Print("enter number  :")
	fmt.Scanf("%d", &Input_number)
	for Counter = 0; Counter < Input_number-2; Counter++ {
		if Counter == 0 {
			sum = 2
			Temperory_Value = 1
		} else {
			First_integer = Second_integer
			Second_integer = Temperory_Value
			Temperory_Value = First_integer + Second_integer
			sum = sum + Temperory_Value

		}

	}
	fmt.Printf("Fibo Series Sum is : %d", sum)

}
