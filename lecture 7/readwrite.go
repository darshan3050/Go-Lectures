package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		detail := recover()
		fmt.Println(detail)

	}()
	// data := []byte("this is sample text for test ")
	file, err := os.Open("text.txt")
	checkError(err)
	file.WriteString("this is sample text for test")
	defer file.Close()
	// data1 := []byte("this is sample text for ioutil test ")
	// err1 := ioutil.WriteFile("text.txt", data1, 0644)
	// checkError(err1)
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
