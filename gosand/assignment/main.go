package main
import "fmt"
func main() {
	a := [3]int{1, 2, 3}
	b := a
	a[1] = 8
	fmt.Print(b)
}
