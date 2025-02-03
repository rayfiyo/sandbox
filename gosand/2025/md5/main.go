package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	text1 := "答え：家とは金魚鉢のことで、男と女は２匹の金魚。飼い主が長期間留守にして餌を与えられず、餓死してしまったのだ。"
	text2 := "家とは金魚鉢のことで、男と女は２匹の金魚。飼い主が長期間留守にして餌を与えられず、餓死してしまったのだ。"

	fmt.Printf("%x\n", md5.Sum([]byte(text1)))
	fmt.Printf("%x\n", md5.Sum([]byte(text2)))
}
