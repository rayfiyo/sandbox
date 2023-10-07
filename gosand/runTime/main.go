package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	for i := 0; i < 500000; i++ {
		fmt.Printf("")
	}
	fmt.Printf("経過: %vms\n", time.Since(now).Milliseconds())
}
