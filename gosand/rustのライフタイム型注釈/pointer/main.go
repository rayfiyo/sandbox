package main

import (
	"fmt"
)

func func1(s1 *string) {
	fmt.Println(*s1)
	*s1 = "q"
	fmt.Println(*s1)
}

func main() {
	var s1 *string
	fmt.Println(*s1)

	func1(s1)
}

/*
func noname(s1 *string, s2 *string, hello *string) {
        world := "World"
        s1 = func1(hello, &world)
        fmt.Println("(1)")
}

func main() {
        var s1 *string
        hello := "Hello"

        noname(s1, s2, &hello)

        fmt.Println("(3)")
        fmt.Println("(4)", *s1)
}
*/
