package factorialOperation

import "fmt"

func calculationMain(n int) int {
	if n > 1 {
		n--
		n = (n + 1) * calculationMain(n)
	}
	return n
}

func UseRecursiveFunc(n int) {
	fmt.Printf("use recursive func: ")
	fmt.Println(calculationMain(n))
}

/*
// https://qiita.com/atsutama/items/667dc25aa3535ffe9e79#functional-go-programmer

package fac

func Factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return Factorial(n - 1) * n
	}
}

*/
