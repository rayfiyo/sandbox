package main

import (
	"fmt"
)

func main() {
	fmt.Println("インプット開始………")
	for i, k := 2, 0; i < 220; i += 2 {
		k++
		if k%5 == 0 {
			i++
		}
		fmt.Println((i-(k/5))/2)
	}
}
