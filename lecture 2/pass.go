package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Enter Number = ")
	fmt.Scanln(&a)
	fmt.Print("Enter Number = ")
	fmt.Scanln(&b)
	swap(&a, &b)

	fmt.Println(a, "  ", b)
}
func swap(a, b *int) {
	var temp *int
	temp = a
	a = b
	b = temp

}
