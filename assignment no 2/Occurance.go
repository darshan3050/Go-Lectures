package main

import "fmt"

func main() {

	var Numner_to_find, No_of_inputs int
	var Array [20]int
	var Count = 0
	fmt.Println("enter no:")
	fmt.Scanf("%d", &No_of_inputs)

	for i := 0; i < No_of_inputs; i++ {

		fmt.Scan(&Array[i])

	}
	fmt.Println("Number to Find :")
	fmt.Scan(&Numner_to_find)

	for i := 0; i <= No_of_inputs; i++ {
		if Array[i] == Numner_to_find {
			Count = Count + 1
		}
	}
	fmt.Println("Occurance is: ", Count)
}
