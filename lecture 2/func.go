package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Enter Number = ")
	fmt.Scanln(&a)
	fmt.Print("Enter Number = ")
	fmt.Scanln(&b)
	c := uni(add, a, b)
	d := uni(sub, a, b)
	e := uni(mul, a, b)
	f := uni(div, a, b)
	fmt.Println(c, " ", d, " ", e, " ", f)
}
func add(x, y int) float32 {
	sum := float32(x) + float32(y)
	return sum
}

func sub(x, y int) float32 {
	sum := float32(x) - float32(y)
	return sum
}

func mul(x, y int) float32 {
	sum := float32(x) * float32(y)
	return sum
}
func div(x, y int) float32 {

	sum := float32(x) / float32(y)
	return sum
}
func uni(funcname func(p, q int) float32, a, b int) float32 {
	ans := funcname(a, b)
	return ans
}
