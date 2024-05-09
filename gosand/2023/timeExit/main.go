package main

import (
	"fmt"
	"time"
)

func main() {
	// 1500ms後に終了するチャネルを作成
	timeLimit := time.Millisecond * 1500
	done := time.After(timeLimit)
	now := time.Now()

	// チャネルが通知されるまで処理を実行
	for {
		select {
		case <-done:
			// チャネルが通知されたら終了
			fmt.Println("終了")
			fmt.Printf("経過: %vms\n", time.Since(now).Milliseconds())
			return
		default:
			// チャネルが通知されるまで処理を実行
			fmt.Printf("")
			continue
		}
	}
}
