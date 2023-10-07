package main

import (
	"fmt"
	"sort"
)

func main() {
	// mapの値を降順に並び替え，キーを出力
	// 3 2 1 4 (2 1 0 3 に 1加えたもの) の出力を期待
	N := map[int]int{0: 3, 1: 3, 2: 7, 3: 1}

	keys := make([]int, 0, len(N))
	for v := range N {
		keys = append(keys, v)
	}
	sort.Slice(keys, func(i, j int) bool {
		return N[keys[i]] > N[keys[j]]
	})

	for _, v := range keys {
		fmt.Println(v + 1)
	}

}
