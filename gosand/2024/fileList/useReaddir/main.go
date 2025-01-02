package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("/")
	files, _ := f.Readdir(0)

	for _, v := range files {
		if v.IsDir() == true {
			fmt.Println(v.Name())
		}
	}
}
