package main

import (
	"fmt"
	"sync"
	"time"
)

var wait = sync.WaitGroup{}

func main() {
	wait.Add(1)
	ch := make(chan int, 2)
	go WriteToChannel(ch)
	fmt.Println("main Print")

	value := <-ch
	value2 := <-ch
	fmt.Println("Value :", value)
	fmt.Println("Value2 :", value2)
	wait.Wait()
}
func WriteToChannel(ch chan int) {
	time.Sleep(time.Second)
	ch <- 20
	ch <- 12
	fmt.Println("Chanel in Writing")
	wait.Done()
}
