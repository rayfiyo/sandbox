package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func floatToBinary(f float64, prec int) string {
	// 整数部を2進数文字列に変換
	intPart := int64(f)
	intStr := strconv.FormatInt(intPart, 2)

	// 小数部を取り出して2進数に変換（指定桁数まで）
	frac := f - float64(intPart)
	fracStr := ""
	for i := 0; i < prec; i++ {
		frac *= 2
		if frac >= 1 {
			fracStr += "1"
			frac -= 1
		} else {
			fracStr += "0"
		}
	}

	return intStr + "." + fracStr
}

// binaryToFloat は "整数部.小数部" の形式の 2 進数文字列を 10 進数の float64 に変換する。
// 小数部は、各桁ごとに 2 の負のべき乗として計算する。
func binaryToFloat(binStr string) (float64, error) {
	parts := strings.Split(binStr, ".")
	if len(parts) != 2 {
		return 0, fmt.Errorf("無効な形式の 2 進数文字列: %s", binStr)
	}

	// 整数部の変換
	intVal, err := strconv.ParseInt(parts[0], 2, 64)
	if err != nil {
		return 0, err
	}
	f := float64(intVal)

	// 小数部の変換
	for i, r := range parts[1] {
		switch r {
		case '1':
			f += math.Pow(2, -(float64(i) + 1))
		case '0':
			// 何もしない
		default:
			return 0, fmt.Errorf("不正な文字 '%c' が含まれています", r)
		}
	}

	return f, nil
}

func main() {
	// 10進数リテラル 0.2（内部表現は IEEE754 による丸めが起こる）
	dec := 0.2
	prec := 60 // 小数部を何桁まで表示するか

	fmt.Printf("10進数 値:  %0.54f\n", dec)

	decBinary := floatToBinary(dec, prec)
	fmt.Printf("2進数表現:  %s\n", decBinary)

	decConverted, err := binaryToFloat("0.001100110011001100110011001100110011001100110011001101")
	if err != nil {
		log.Fatal("2進数から10進数への変換エラー：", err)
	}
	fmt.Printf("逆変換 値:  %0.60f\n", decConverted)

	decConverted, err = binaryToFloat("0.001100110011001100110011001100110011001100110011001100")
	if err != nil {
		log.Fatal("2進数から10進数への変換エラー：", err)
	}
	fmt.Printf("逆変換 値:  %0.60f\n", decConverted)
}
