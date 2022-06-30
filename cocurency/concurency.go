package main

import (
	"fmt"
	"sync"
	"time"
)

var wait_group = sync.WaitGroup{}

func main() {
	wait_group.Add(2)
	go delay1()
	go delay2()
	wait_group.Wait()
}
func delay1() {
	for i := 0; i < 7; i++ {
		{
			time.Sleep(time.Second)
			fmt.Println("delay 1 func index: ", i)
		}
	}
	wait_group.Done()
}
func delay2() {
	for i := 0; i < 5; i++ {
		{
			time.Sleep(time.Second)
			fmt.Println("delay 2 func index: ", i)
		}
	}
	wait_group.Done()
}
