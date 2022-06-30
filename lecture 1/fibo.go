package main

import "fmt"

func main() {
	var a = 0
	var b = 1
	var sum, c, n, t int
	fmt.Print("enter number")
	fmt.Scanf("%d", &n)
	for c = 0; c < n-2; c++ {
		if c == 0 {
			sum = 2
			t = 1
		} else {
			a = b
			b = t
			t = a + b
			sum = sum + t

		}

	}
	fmt.Printf("%d", sum)

}
