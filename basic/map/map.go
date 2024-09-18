package main

import "fmt"

func main() {
	m := make(map[int]int)
	go func() {
		m[1] = 1
	}()
	m[2] = 2
	fmt.Println(m[1])
}
