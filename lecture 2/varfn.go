package main

import "fmt"

func main() {
	var s []string
	fmt.Scanln(&s)
	name(s)
}
func name(name ...string) {
	for _, v := range name {
		fmt.Println(v, " ")
	}
}
