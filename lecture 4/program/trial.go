package program

import "fmt"

func Prime(Input_number int) {
	var Counter, Count int
	fmt.Scanf("%d", &Input_number)

	if Input_number < 2 {
		println("invalid number")

	}
	for Counter = 2; Counter < Input_number/2; Counter++ {
		if Input_number%Counter == 0 {
			Count = 1
		}
	}
	if Count == 1 {
		fmt.Println("Not Prime Number")
	} else {
		fmt.Println("Prime Number")
	}

}
