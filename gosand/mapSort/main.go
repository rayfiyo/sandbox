package main

import (
	"fmt"
	"sort"
)

func main() {
	// mapの値を降順に並び替え，キーを出力
	// 3 2 1 4 (2 1 0 3 に 1加えたもの) の出力を期待
	N := map[int]int{0: 3, 1: 3, 2: 7, 3: 1}
	ans := sortmap(N)
	fmt.Println(ans)
}

func sortmap(N map[int]int) []int {
	/*
		// mapの値を降順に並び替え，キーを出力
		// 3 2 1 4 (2 1 0 3 に 1加えたもの) の出力を期待
		N := map[int]int{0: 3, 1: 3, 2: 7, 3: 1}
		// mapのキーをkeysスライスにコピー
	*/
	keys := make([]int, 0, len(N))
	for vofN := range N { // Nのキー
		keys = append(keys, vofN)
	}
	sort.Slice(keys, func(i, j int) bool { // Nに基づき並び替え
		if N[keys[i]] != N[keys[j]] {
			return N[keys[i]] > N[keys[j]] // 値の降順に基づく
		} else {
			return keys[i] < keys[j] // 同値のときキーの昇順に基づく
		}
	})
	return keys
	/*
		for _, v := range keys {
			fmt.Println(v + 1)
		}
	*/
}

/*
   	// キーをスライスにコピー
   	keys := make([]int, 0, len(values))
   	for key := range values {
   		keys = append(keys, key)
   	}

   	// スライスを値に基づいて降順にソート
   	sort.Slice(keys, func(i, j int) bool {
   		return values[keys[i]] > values[keys[j]]
   	})

   	// ソートされたキーを元にキーと値を表示
   	for _, key := range keys {
   		fmt.Printf("i=%dのとき %d\n", key, values[key])
   	}
   }
*/

/*
   func main() {
   	// 与えられた値をマップに格納
   	values := map[int]int{
   		0: 3,
   		1: 3,
   		2: 7,
   		3: 1,
   	}

   	// キーをスライスにコピー
   	keys := make([]int, len(values))
   	i := 0
   	for key := range values {
   		keys[i] = key
   		i++
   	}

   	// スライスを値に基づいて降順にソート
   	sort.Slice(keys, func(i, j int) bool {
   		return values[keys[i]] > values[keys[j]]
   	})

   	// ソートされたキーを元にキーと値を表示
   	for _, key := range keys {
   		fmt.Printf("i=%dのとき %d\n", key, values[key])
   	}
   }
*/
