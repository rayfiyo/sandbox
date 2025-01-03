package main

import (
	"fmt"
)

func main() {
	s := "ahoge"
	fmt.Scanln(&s)

	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	fmt.Println(string(r))
}
