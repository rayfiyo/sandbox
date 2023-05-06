package factorialOperation

import "fmt"

func UseFor(n int) {
	fmt.Printf("use for: ")
	sum := 1
	for i := 1; i <= n; i++ {
		sum *= i
	}
	fmt.Println(sum)
}

/*
// https://qiita.com/atsutama/items/667dc25aa3535ffe9e79#junior-go-programmer

package fac

func Factorial(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

*/
