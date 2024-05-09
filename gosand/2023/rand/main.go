package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	fmt.Println(rand.Intn(1))
	fmt.Println(rand.Intn(30))
	fmt.Println(rand.Intn(30))
	fmt.Println(rand.Intn(30))
}
