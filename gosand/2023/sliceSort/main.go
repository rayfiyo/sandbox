package main

import (
	"sort"
)

func main() {
	a := make([]int, 5)
	a[0], a[1], a[2], a[3], a[4] = 1, 3, 2, 5, 4
	sort.Ints(a)
}
