package main

import "fmt"

const (
	a = iota + 1
	_ // 右辺の省略なので、当然ここでも iota は ++
	b // 右辺の省略
	c = iota * 2
)

func main() {
	fmt.Println(a, b, c)
}
