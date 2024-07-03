package main

import "fmt"

func a(data chan int) {
	fmt.Println(data)
}

func main() {
	c := make(chan int)

	a(c)
}
