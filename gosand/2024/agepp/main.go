package main

import (
	"fmt"
	"time"
)

func main() {
	age := 18
	t := time.Now()

	if t.Month() == time.August && t.Day() == 27 {
		age++
	}

	fmt.Println(age)
}
