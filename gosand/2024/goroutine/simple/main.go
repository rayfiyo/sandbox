// 簡素な goroutine を使ったコード

package main

import (
	"fmt"
	"time"
)

func sendMessage(msg string) {
	fmt.Println(msg)
}

func main() {
	message := "hi"
	go sendMessage(message)
	message = "bye"

	// time.Sleep(time.Second) // この行をコメントアウトすると bye が先
	fmt.Println(message)
	time.Sleep(time.Second)
}
