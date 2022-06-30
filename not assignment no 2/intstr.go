package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a = 10
	b := strconv.Itoa(a)

	fmt.Printf("%T   %v ", b, b)
}
