package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)
// バッファでstrings型として行読み取り
func stringLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	S := stringLine()

	fmt.Println(string(S[99]))
}
