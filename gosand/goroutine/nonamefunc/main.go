// 無名関数では，実行のタイミングがことなる

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
	go func() {
		sendMessage(message) // hi の出力を期待するが， bye が出力される
	}()
	message = "bye"

	time.Sleep(time.Second)
	fmt.Println(message)
	time.Sleep(time.Second)
}
