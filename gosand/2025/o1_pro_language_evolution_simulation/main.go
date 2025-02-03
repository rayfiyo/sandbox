package main

import (
	"fmt"
	"math/rand"
	"time"
)

// CulturalState: ある仮想コミュニティの「言語・文化状態」を持つ構造体
type CulturalState struct {
	Population    int
	CommonWords   []string // 代表的によく使われる単語
	UniqueCustoms []string // そのコミュニティ独特の風習
	LastEvent     string
}

// LLMStub: LLMを利用するためのダミーインタフェース
type LLMStub struct{}

func (l *LLMStub) GenerateCultureChange(state CulturalState, userInputs []string) CulturalState {
	// 実際にはLLM APIを呼び出し、今の状態やユーザの入力(コミュニケーション内容)をpromptに含めて
	// 「次に起こりうる文化的変化」や「新しい単語」の創出をさせる

	// ここではダミー実装
	newState := state
	// ユーザが使った単語の一部を文化に取り込む(ランダムに)
	if len(userInputs) > 0 {
		pick := userInputs[rand.Intn(len(userInputs))]
		newState.CommonWords = append(newState.CommonWords, pick)
	}
	// ランダムイベントを生成
	newState.LastEvent = "異文化との接触があり、新たな慣習が生まれたかもしれない"
	return newState
}

// シミュレーションを1サイクル進める
func simulateStep(state CulturalState, llm *LLMStub, userInputs []string) CulturalState {
	// (1) LLMを使って次の状態を決定
	newState := llm.GenerateCultureChange(state, userInputs)
	// (2) 簡単なルール: 人口の増減(ダミー)
	newState.Population += rand.Intn(5) - 2
	if newState.Population < 0 {
		newState.Population = 0
	}
	return newState
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// 初期のコミュニティ設定
	state := CulturalState{
		Population:    100,
		CommonWords:   []string{"hello", "world"},
		UniqueCustoms: []string{"danceFestival"},
		LastEvent:     "特に大きな出来事はない",
	}

	llm := &LLMStub{}

	// ユーザがチャットや投稿で使った単語 (擬似)
	userInputs := []string{"kawaii", "arigato", "cool"}

	// 5サイクルほどシミュレーション
	for i := 0; i < 5; i++ {
		state = simulateStep(state, llm, userInputs)

		fmt.Printf("---- Cycle %d ----\n", i+1)
		fmt.Printf("Population: %d\n", state.Population)
		fmt.Printf("CommonWords: %v\n", state.CommonWords)
		fmt.Printf("UniqueCustoms: %v\n", state.UniqueCustoms)
		fmt.Printf("LastEvent: %s\n\n", state.LastEvent)
	}
}
