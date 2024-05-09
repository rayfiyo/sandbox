// AtCoderのコンテストである、 AtCoder Beginner Contest。
// これの コンテスト全番号001~297 (2023年5月9日現在) のうち、11の最大約数(11で割れる最大値)を求める。
package main

import "fmt"

func main() {
	fmt.Printf("ABC001～297で、11の最大約数:")

	for i := 297; i > 0; i-- {
		if i%11 == 0 {
			fmt.Print(i)
			i = 0	//break
		}
	}
}
