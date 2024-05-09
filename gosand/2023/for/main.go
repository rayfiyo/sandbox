package main

import (
	"fmt"
)

func main() {
	N, A, B, C, sum := 300, 1000, 1000, 0, 2000
	fmt.Scanf("%d %d %d", &N, &A, &B)
	sum = A + B
	fmt.Scanf("%d", &C)
	for i := 0; i < N; i++ {
		fmt.Println(i)
		if C == sum {
			fmt.Println("true")
		} else {
			fmt.Println("lnlnlnll")
			fmt.Scanf("%d", &C)
		}
	}
}
