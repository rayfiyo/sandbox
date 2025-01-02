package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printSlowly(s string, delay time.Duration) {
	for _, char := range s {
		fmt.Print(string(char))
		time.Sleep(delay)
	}
	fmt.Println()
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	printSlowly("🎉🎂 特別な誕生日プログラムを Claude がプレゼントします！ 🎂🎉", 80*time.Millisecond)

	messages := []string{
		// "誕生日おめでとう！年齢を重ねるごとにバグは減っていくはず…たぶん。",
		// "またひとつ賢くなりましたね。でも、コンパイルエラーは相変わらずですよ？",
		// "Happy Birthday! 今年こそはStackOverflowなしでコードが書けますように。",
		"お誕生日おめでとう！あなたのライフサイクルにまた1年がコミットされました。",
		// "誕生日だからって、セミコロンを忘れちゃダメですよ！",
	}

	selectedMessage := messages[rand.Intn(len(messages))]
	printSlowly(selectedMessage, 80*time.Millisecond)

	cake := `
	    ,,,,,
	   .(◕‿◕).
	  /-------\
	 /  *   *  \
	|  * 🎂  *  |
	|__*_____*__|
	`
	fmt.Println(cake)

	printSlowly("さぁ、ロウソクを消しましょう！", 80*time.Millisecond)
	for i := 3; i > 0; i-- {
		printSlowly(fmt.Sprintf("%d...", i), 80*time.Millisecond)
		time.Sleep(800 * time.Millisecond)
	}
	printSlowly("フーッ！💨", 80*time.Millisecond)

	printSlowly("おめでとうございます！素敵な1年になりますように！！", 80*time.Millisecond)
}
