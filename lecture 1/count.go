package main

import "fmt"

func main() {
	var a [20]int
	var n, i int
	fmt.Println("enter no:")
	fmt.Scanf("%d", &n)
	for i = 1; i < n; i++ {

		fmt.Scanf("  %d  ", &a[i])

	}

	for i = 0; i < n; i++ {
		if a[i]%2 != 0 {
			fmt.Printf("/n %d is odd", a[i])
		} else if a[i]%2 == 0 {
			fmt.Printf("/n %d is even", a[i])
		} else {
			fmt.Printf("/n %d is zero", a[i])
		}

	}
}
