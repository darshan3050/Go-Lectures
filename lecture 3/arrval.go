package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	val(a)
	ref(&a)
	fmt.Println(a[:])

}
func val(arr [5]int) {
	fmt.Println(arr)
	arr[0] = 300
	fmt.Println(arr)
}
func ref(arr *[5]int) {
	fmt.Println(arr)
	arr[0] = 500
	fmt.Println(arr)
}
