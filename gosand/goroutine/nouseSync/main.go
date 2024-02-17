// main関数が goroutine の終了を待たない（syncパッケージを使わない）

package main

import (
	"fmt"
	"time"
)

func heavyMessage() {
	time.Sleep(time.Second)
	fmt.Println("wait!")
}

func main() {
	fmt.Println("hi")

	go func() {
		heavyMessage()
	}()

	fmt.Println("bye")
}
