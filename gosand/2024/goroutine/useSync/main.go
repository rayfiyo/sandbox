// main関数が goroutine の終了を待つ（syncパッケージを使う）

package main

import (
	"fmt"
	"sync"
	"time"
)

func heavyMessage() {
	time.Sleep(time.Second)
	fmt.Println("wait!")
}

func main() {
	fmt.Println("hi")

	var wg sync.WaitGroup
	wg.Add(1) // リファレンスカウンタを +1 する
	go func() {
		defer wg.Done() // リファレンスカウンタを -1 する
		heavyMessage()
	}()

	fmt.Println("bye")

	wg.Wait() // リファレンスカウンタが0になるまで待つ
}
