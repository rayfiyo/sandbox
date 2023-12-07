package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	timeLimit := time.Millisecond * 1990

	// コンテキストの作成
	ctx, cancel := context.WithTimeout(context.Background(), timeLimit)
	defer cancel()

	// 非同期で処理を実行
	now := time.Now()
	go func() {
		// 処理を書く
		for {
			fmt.Printf("")
		}
	}()

	// コンテキストを待機
	select {
	case <-ctx.Done():
		// タイムアウトやキャンセルが発生した場合の処理
		fmt.Printf("経過: %vms\n", time.Since(now).Milliseconds())
		fmt.Println("Timeout or canceled:", ctx.Err())
	}
}
