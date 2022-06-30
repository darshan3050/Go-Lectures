package main

import "fmt"

func main() {
	var Array [20]int
	var No_of_inputs, Counter int
	Zero_counter := 0
	Even_counter := 0
	Odd_counter := 0
	fmt.Println("enter no:")
	fmt.Scanf("%d", &No_of_inputs)

	for Counter = 0; Counter < No_of_inputs; Counter++ {

		fmt.Scan(&Array[Counter])

	}

	for Counter = 0; Counter < No_of_inputs; Counter++ {
		if Array[Counter] == 0 {

			Zero_counter = Zero_counter + 1
		} else if Array[Counter]%2 == 0 {

			Even_counter = Even_counter + 1
		} else {

			Odd_counter = Odd_counter + 1
		}

	}
	fmt.Println("Number of Even :", Even_counter)
	fmt.Println("Number of Odd :", Odd_counter)
	fmt.Println("Number of Zero :", Zero_counter)
}
