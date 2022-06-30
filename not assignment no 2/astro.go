package main

import "fmt"

func main() {
	var b int

	var p, q, r, sum int
	for i := 100; i < 1000; i++ {
		b = i
		p = b % 10
		b = b / 10
		q = b % 10
		b = b / 10
		r = b
		b = i
		sum = p*p*p + q*q*q + r*r*r

		if sum == i {
			fmt.Printf("no %d\n", b)
		}

	}

}
