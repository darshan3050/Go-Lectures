package main

import "fmt"

func main() {
	var i int
	var Array [20]int
	var Answer_array [20]int
	var No_of_inputs, Counter int
	fmt.Println("enter no:")
	fmt.Scanf("%d", &No_of_inputs)

	for Counter = 0; Counter < No_of_inputs; Counter++ {

		fmt.Scan(&Array[Counter])

	}
	fmt.Println("first array: ")
	for i = 0; i < No_of_inputs; i++ {
		fmt.Printf("%d   ", Array[i])
	}
	Answer_array = sort(Array, No_of_inputs)

	fmt.Println("\nSorted Array: ")
	for i = 0; i < No_of_inputs; i++ {
		fmt.Printf("%d   ", Answer_array[i])
	}
}
func sort(Function_array [20]int, Total_no int) [20]int {
	var i, j, Temperory_variable int
	fmt.Println("Inside function\n ")
	for i = 0; i < Total_no-1; i++ {
		for j = 0; j < Total_no-1; j++ {
			if Function_array[j] > Function_array[j+1] {
				Temperory_variable = Function_array[j]
				Function_array[j] = Function_array[j+1]
				Function_array[j+1] = Temperory_variable
			}
		}
	}
	fmt.Println("Function array print : ")
	for i = 0; i < Total_no; i++ {
		fmt.Printf("%d   ", Function_array[i])
	}
	return Function_array
}
