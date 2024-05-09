package main

import (
	"fmt"
	"github.com/YosCiiCable/sandbox/gosand/factorial/cmd"
)

func main() {
	n := 0
	fmt.Printf("n階乗のn: ")
	fmt.Scanf("%d", &n)

	factorialOperation.UseFor(n)
	factorialOperation.UseRecursiveFunc(n)
}
