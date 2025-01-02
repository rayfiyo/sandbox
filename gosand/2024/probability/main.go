// 8択の内2択は再度選択するときの確率を求めるプログラム

package main

import (
	"fmt"
	"math/rand/v2"
)

const (
	// choices の数（選択肢の数）．len(choices)
	N = 10

	// 選択をやり直す選択肢
	R1 = 1
	R2 = 5
	R3 = 8
)

func dice(c *[N]int) {
	n := rand.N(N)

	switch n {
	case R1:
		dice(c)
	case R2:
		dice(c)
	case R3:
		dice(c)
	default:
		c[n]++
	}
}

func printSel(c *[N]int, n int) {
	if n%12000000 == 0 {
		fmt.Printf("試行%24d: {", n)

		n := float64(n)

		for i, v := range *c {
			if i != N-1 {
				fmt.Printf("%2.4f%%, ", float64(v)/n*100)
			} else {
				fmt.Printf("%2.4f%% }\n", float64(v)/n*100)
			}
		}
	}
}

func main() {
	choices := [N]int{0, 0, 0, 0, 0, 0, 0, 0}

	fmt.Println(len(choices), " つの選択肢")

	for i := 1; ; i++ {
		dice(&choices)

		printSel(&choices, i)
	}
}
