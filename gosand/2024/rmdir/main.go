package main

import "fmt"
import "os"

func main() {
    err := os.Remove("/home/ray/go/src/github.com/rayfiyo/sandbox/gosand/rmdir/test")
    fmt.Println(err)
}
