package main

import "fmt"

func main() {
	var Largest_no int = 0
	var Second_largest_no int = 0
	var Array [20]int
	var No_of_inputs, Counter int
	fmt.Println("enter no:")
	fmt.Scanf("%d", &No_of_inputs)

	for Counter = 0; Counter < No_of_inputs; Counter++ {

		fmt.Scan(&Array[Counter])

	}

	Largest_no = Array[0]
	for i := 1; i <= 4; i++ {
		if Largest_no < Array[i] {
			Second_largest_no = Largest_no
			Largest_no = Array[i]
		} else if Second_largest_no < Array[i] {
			Second_largest_no = Array[i]
		}
	}
	fmt.Println("Second largest element is: ", Second_largest_no)
}
