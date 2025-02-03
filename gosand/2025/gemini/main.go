package main

import (
	"context"
	"fmt"
	"log"
	"os"

	// ここは実際のパッケージに合わせて import 名を変えてください
	// 仮に "cloud.google.com/go/genai" のような公式ライブラリがある前提とします
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type GeminiLLMGateway struct {
	client *genai.Client
	model  *genai.GenerativeModel
}

// NewGeminiLLMGateway: コンストラクタ

func NewGeminiLLMGateway(apiKey string) (*GeminiLLMGateway, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY is not set")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create gemini client: %w", err)
	}

	model := client.GenerativeModel("gemini-2.0-flash-exp")

	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text(`このコミュニティの文化を新しい方向に進化させるアイデアを提案してください。
    あなたの出力は**必ず次のJSON形式**で返してください。
    {
        "newCulture": "string",
        "populationChange": 0
    }`),
		},
	}

	return &GeminiLLMGateway{
		client: client,
		model:  model,
	}, nil
}

// GenerateCultureUpdate: LLMGatewayインタフェース
// ここではストリーミングなしで結果をまとめて取得する例
func (g *GeminiLLMGateway) GenerateCultureUpdate(prompt string) string {
	ctx := context.Background()
	resp, err := g.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintln(resp.Candidates[0].Content.Parts)
}

// GenerateCultureUpdate: LLMGatewayインタフェース
// ここではストリーミングありで結果をまとめて取得する例
func (g *GeminiLLMGateway) GenerateCultureUpdateStreaming(ctx context.Context, prompt string) string {
	iter := g.model.GenerateContentStream(ctx, genai.Text(prompt))
	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		_ = resp /* 逐次出力（ストリーム）
		fmt.Println(resp.Candidates[0].Content.Parts)
		// */
	}

	return fmt.Sprintln(iter.MergedResponse().Candidates[0].Content.Parts)
}

func main() {
	apiKey := os.Args[1]
	g, err := NewGeminiLLMGateway(apiKey)
	if err != nil {
		log.Fatal(err)
	}
	defer g.client.Close()

	fmt.Println(g.GenerateCultureUpdate(`踊りの文化`))
	// fmt.Println(g.GenerateCultureUpdateStreaming(ctx, `踊りの文化`))
}
