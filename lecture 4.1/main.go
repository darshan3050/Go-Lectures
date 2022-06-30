package main

import (
	"fmt"
	"go/statistic"
)

func main() {
	var Array []float64
	var No_of_elements, i int
	var Mean_answer, Median_answer, Mode_answer, Temperory float64
	fmt.Println("Enter total no of elements: ")
	fmt.Scan(&No_of_elements)
	for i = 1; i < No_of_elements; i++ {

		fmt.Scan("  %d  ", &Temperory)
		Array = append(Array, Temperory)

	}

	Mean_answer = statistic.Mean_Function(Array)
	Median_answer = statistic.Mean_Function(Array)
	Mode_answer = statistic.Mean_Function(Array)

	println("Mean Value :")
	println(Mean_answer)
	println("Median Value :")
	println(Median_answer)
	println("Mode Value :")
	println(Mode_answer)
}
