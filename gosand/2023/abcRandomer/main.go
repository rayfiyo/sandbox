package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	randNum := rand.Intn(10000) + 10
	for i := 0; i < randNum; i++ {
		rand.NewSource(time.Now().UnixNano())
		switch rand.Intn(3) {
		case 0:
			fmt.Printf("A")
		case 1:
			fmt.Printf("B")
		case 2:
			fmt.Printf("C")
		}
	}
	fmt.Println()
}
