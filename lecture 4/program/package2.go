package program

import "fmt"

func Astro() {
	var Temperory_Value int

	var First_Seprated_Integer, Second_Seprated_Integer, Third_Seprated_Integer, Astronomical_sum int
	for i := 100; i < 1000; i++ {
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
