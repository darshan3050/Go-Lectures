package main

import "fmt"

func main() {
	var Temperory_Value int
	var Astro_minimum_value, Astro_maximum_value int
	fmt.Println("Enter Minimum Range")
	fmt.Scan(&Astro_minimum_value)
	fmt.Println("Enter Maximum Range")
	fmt.Scan(&Astro_maximum_value)
	var First_Seprated_Integer, Second_Seprated_Integer, Third_Seprated_Integer, Astronomical_sum int
	for i := Astro_minimum_value; i < Astro_maximum_value; i++ {
		Temperory_Value = i
		First_Seprated_Integer = Temperory_Value % 10
		Temperory_Value = Temperory_Value / 10
		Second_Seprated_Integer = Temperory_Value % 10
		Temperory_Value = Temperory_Value / 10
		Third_Seprated_Integer = Temperory_Value
		Temperory_Value = i
		Astronomical_sum = First_Seprated_Integer*First_Seprated_Integer*First_Seprated_Integer + Second_Seprated_Integer*Second_Seprated_Integer*Second_Seprated_Integer + Third_Seprated_Integer*Third_Seprated_Integer*Third_Seprated_Integer

		if Astronomical_sum == i {
			fmt.Printf("no %d\n", Temperory_Value)
		}

	}

}
