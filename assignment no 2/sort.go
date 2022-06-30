package main

import "fmt"

func main() {
	var i, No_of_inputs int

	var Array = [5]int{1, 2, 3, 4, 5}
	var Second_array [5]int

	fmt.Println("enter no if inputs :")
	fmt.Scanf("%d", &No_of_inputs)

	fmt.Println("Enter the Values ")
	for i = 0; i < No_of_inputs; i++ {

		fmt.Scan(&Array[i])

	}
	fmt.Println("first array: ")
	for i = 0; i < No_of_inputs; i++ {
		fmt.Printf("%d   ", Second_array[i])
	}

	Second_array = Sort(Array)

	fmt.Println("Sorted Array:\n ")
	for i = 0; i < No_of_inputs; i++ {
		fmt.Printf("%d   ", Second_array[i])
	}
	println()
}

func Sort(Function_array [5]int) []int {

	var i, j, Temperory_variable, Total_no int
	fmt.Println("Inside function\n ")
	for i = 0; i < Total_no; i++ {
		for j = 0; j < Total_no; j++ {
			if Function_array[j] > Function_array[j+1] {
				Temperory_variable = Function_array[j]
				Function_array[j] = Function_array[j+1]
				Function_array[j+1] = Temperory_variable
			}
		}
	}
	fmt.Println("Sorted Array inside function:\n ")
	for i = 0; i < Total_no; i++ {
		fmt.Printf("%d   ", Function_array[i])
	}
	println()
	return Function_array[:]
}
