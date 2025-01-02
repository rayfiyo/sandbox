package main

import (
	"fmt"
)

func func1(hello *string, world *string) *string {
	fmt.Println("func1", *hello, *world)
	return hello
}

func func2(hello *string, world *string) *string {
	fmt.Println("func2", *hello, *world)
	return world
}

func noname(s1 *string, s2 *string, hello *string) {
	world := "World"
	_ = s1
	_ = s2
	s1 = func1(hello, &world)
	fmt.Println("(1)")
	s2 = func2(hello, &world)
	fmt.Println("(2)")
	_ = s1
	_ = s2
}

func main() {
	var s1 *string
	var s2 *string
	hello := "Hello"

	/*
		func() {
			world := "World"
			s1 = func1(&hello, &world)
			fmt.Println("(1)")
			s2 = func2(&hello, &world)
			fmt.Println("(2)")
		}()

		fmt.Println("(3)")
		fmt.Println(*s1)
		fmt.Println("(4)")
		fmt.Println(*s2)
		fmt.Println("(5)")
		// */

	// /*
	noname(s1, s2, &hello)

	fmt.Println("(3)")
	fmt.Println("(4)", *s1)
	fmt.Println("(5)", *s2)
	// */
}
