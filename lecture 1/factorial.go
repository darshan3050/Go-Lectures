package main

import "fmt"

func main() {

	var n, ans int
	ans = 1

	fmt.Print("Enter Number = ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		ans = ans * i
	}
	fmt.Printf("The Factorial is   = %d", ans)
}
