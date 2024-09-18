package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("i am goroutine")
		ch <- 1
	}()
	fmt.Println("bye", <-ch)
	fmt.Println("main")
}
