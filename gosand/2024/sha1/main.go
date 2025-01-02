package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"sync"
)

func main() {
	str := "cpaw"
	c := make(chan string, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		sha1 := sha1.New()
		io.WriteString(sha1, str)

		if hex.EncodeToString(sha1.Sum(nil)) == "9216ddf84851f15a46662eb04759d2bebacac666" {
			c <- str
		} else {
			c <- ""
		}
		wg.Done()
	}()

	if v := <-c; v != "" {
		fmt.Println(v)
	}
	wg.Wait()
}
