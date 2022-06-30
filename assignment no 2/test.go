package main

import "fmt"

func main() {
	var i, Temperory_variable, No_of_inputs int

	var Array = make([]int, 10, 10)
	var Second_array = make([]int, 10, 10)

	fmt.Println("enter no if inputs :")
	fmt.Scanf("%d", &No_of_inputs)

	fmt.Println("Enter the Values ")
	for i = 0; i < No_of_inputs; i++ {

		fmt.Scan(&Temperory_variable)
		Array[i] = append(Array[i], Temperory_variable)

	}
	fmt.Println("first array: ")
	for i = 0; i < No_of_inputs; i++ {
		fmt.Printf("%d   ", Second_array[i])
	}

	Second_array = Sort(Array, No_of_inputs)

	fmt.Println("Sorted Array:\n ")
	for i = 0; i < No_of_inputs; i++ {
		fmt.Printf("%d   ", Second_array[i])
	}
	println()
}

func Sort(Array []int, Total_no int) []int {
	Array = make([]int, 10, 10)
	var i, j, Temperory_variable int
	fmt.Println("Inside function\n ")
	for i = 0; i < Total_no; i++ {
		for j = 0; j < Total_no; j++ {
			if Array[j] > Array[j+1] {
				Temperory_variable = Array[j]
				Array[j] = Array[j+1]
				Array[j+1] = Temperory_variable
			}
		}
	}
	fmt.Println("Sorted Array inside function:\n ")
	for i = 0; i < Total_no; i++ {
		fmt.Printf("%d   ", Array[i])
	}
	println()
	return Array
}
