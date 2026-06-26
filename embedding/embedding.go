package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	//model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
	//	APIKey: os.Getenv("ARK_API_KEY"),
	//	Model:  os.Getenv("MODEL"),
	//})

	apiType := ark.APITypeMultiModal
	embedder, err := ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
		APIKey:  os.Getenv("ARK_API_KEY"),
		Model:   os.Getenv("EMBEDDER"),
		APIType: &apiType,
	})
	if err != nil {
		panic(err)
	}
	input := []string{
		"你好，泥嚎",
		"明日方舟",
		"原神",
	}
	embeddings, err := embedder.EmbedStrings(ctx, input)

	if err != nil {
		panic(err)
	}
	//for i, embedding := range embeddings {
	//	//println("文本", i+1, "的向量维度：", len(embedding))
	//	println("文本", i+1, "的向量维度：", embedding)
	//}
	print(embeddings)
	fmt.Println(embeddings)

}
