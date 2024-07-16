package main

import (
	"fmt"

	"github.com/rayfiyo/sandbox/gosand/immutable/entity"
)

func main() {
	money := entity.NewMoney(0)
	user := entity.NewUser("山田", money)

	user.Money = entity.NewMoney(1) // userのmoneyフィールドを新しいMoneyオブジェクトで"交換"する（不変性を守る）

	// 以下は amount にアクセスできないのでコンパイルエラーとなる
	// user.Money.amount = 2 // userが参照するMoneyオブジェクトのamountを直接変更する（不変性に反する）

	// amount が欲しいときは以下を使えば良い.ここでは 1 が出力される
	fmt.Println(user.Money.Amount())
}
